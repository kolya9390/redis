package repository

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/jmoiron/sqlx"
)

type GeoRepository interface {
	Add(query, region, geoLat, geoLon string) error // Вставка в DB's
	Get(query string) ([]AddressData, error) // Получаем данные из базы или из Редиса
	CheckAvailability(query string) (bool, error) // Проверка наличая в базе and Cache
}

type GeoRepositoryDB interface {
	InsertSearchHistory(query string) (int, error) // Вставка строки поиска в Таблицу search_history и надо б кэш
	InsertAddress(region, geoLat, geoLon string) (int, error) // Вставка адреса, и координат в Таблицу address и кэш
	InsertHistorySearchAddress(searchHistoryID, addressID int) error // Вставка айд в Таблицу history_search_address
	SearchInHistory(query string) (bool, error) // Проверка наличая в базе 
	FindAddressByQueryAndHistory(query string) ([]AddressData, error) // Селекс по двум таблицам 

}

type Cacher interface {
    Set(key string, value []AddressData) error // Устанавливает запись в редис
    Get(key string) ([]AddressData, error) // получаем данные из редиса
	Check(query string) (bool, error)
}

type Cache struct {
	client *redis.Client
}


type geoRepository struct {
	db *sqlx.DB
}

type GeoProxy struct {
	geoRepo geoRepository
	cache 	Cacher
}




func NewGeoRepositoryDB(db *sqlx.DB) *geoRepository {
	return &geoRepository{db: db}
}

func NewGeoRedis(client *redis.Client) Cacher {
	return &Cache{
		client: client}
}

func NewGeoRepositoryProxy(repo geoRepository,cache Cacher) *GeoProxy {


    return &GeoProxy{
        geoRepo: repo,
        cache:      cache,
    }
}

func (gp *GeoProxy) Add(query, region, geoLat, geoLon string) error {

	queryID ,err := gp.geoRepo.InsertSearchHistory(query)

	if err!= nil {
		return err
	}

	adressID, err := gp.geoRepo.InsertAddress(region,geoLat,geoLon)

	if err!= nil {
		return err
	}

	err = gp.geoRepo.InsertHistorySearchAddress(queryID,adressID)

	return err
}

func (gp *GeoProxy) Get(query string) ([]AddressData, error) {


	if ok, err := gp.cache.Check(query) ; ok {

		if err!= nil {
			return nil,fmt.Errorf("Check:%s",err)
		}

		resp, err := gp.cache.Get(query)
		if err!= nil {
			return nil,fmt.Errorf("Get in cache : %s",err)
		}

		return resp,nil
	}

	resp,err := gp.geoRepo.FindAddressByQueryAndHistory(query)
	var id int

	if err != nil {
		return nil,fmt.Errorf("FindAddressByQueryAndHistory:%s",err)
	}
	id++

	err = gp.cache.Set(fmt.Sprintf("%s:%s",query,id),resp)

	if err != nil {
		return nil, fmt.Errorf("Set:%s",err)
	}

	return resp,nil
}

func (gp *GeoProxy) CheckAvailability(query string) (bool, error) {

	if ok, err := gp.cache.Check(query) ; ok{

		if err != nil{
			return false,err
		}

		return true,nil
	}

	if OK, err := gp.geoRepo.SearchInHistory(query) ; OK {
		if err != nil {
			return false,err
		}

		return true , nil
	}

	return false, nil

}



func (c *Cache) Set(key string, value []AddressData) error {
	jsonValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return c.client.Set(key, jsonValue, 0).Err()
}

func (c *Cache) Get(key string) ([]AddressData, error) {
	val, err := c.client.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, errors.New(fmt.Sprintf("not found by key %s", key))
		}
		return nil, err
	}
	var value []AddressData
	err = json.Unmarshal([]byte(val), &value)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (c *Cache) Check(query string) (bool, error) {

	exists , err := c.client.Exists(query).Result()

	if err != nil {
		return false ,err
	}


	if exists != 0 {
		return true, redis.Nil
	} 

	return false,nil
}

// создание 3 таблиц
func (d *geoRepository) ConnectToDB() error {


	sqlStatementSearch_history := `
CREATE TABLE IF NOT EXISTS search_history (
    id SERIAL PRIMARY KEY,
    query text
);`

	sqlStatementAddress := `
CREATE TABLE IF NOT EXISTS address (
    id SERIAL PRIMARY KEY,
    region text,
    geo_lat text,
    geo_lon text
);`

	sqlStatementHistory_search_address := `
CREATE TABLE IF NOT EXISTS history_search_address (
    id SERIAL PRIMARY KEY,
    search_history_id int,
    address_id int
);`

setTrgm := `CREATE EXTENSION IF NOT EXISTS pg_trgm;`


	_, err := d.db.Exec(sqlStatementSearch_history)
	if err != nil {
		return err
	}

	_, err = d.db.Exec(sqlStatementAddress)
	if err != nil {
		return err
	}

	_, err = d.db.Exec(sqlStatementHistory_search_address)
	if err != nil {
		return err
	}


	_, err = d.db.Exec(setTrgm)
	if err != nil {
		return err
	}


	return nil

}



func (gr *geoRepository) InsertSearchHistory(query string) (int, error) {
	var id int
	err := gr.db.QueryRowx("INSERT INTO search_history (query) VALUES ($1) RETURNING id", query).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}


func (gr *geoRepository) InsertAddress(region, geoLat, geoLon string) (int, error) {
	var id int
	err := gr.db.QueryRowx("INSERT INTO address (region, geo_lat, geo_lon) VALUES ($1, $2, $3) RETURNING id", region, geoLat, geoLon).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}


func (gr *geoRepository) InsertHistorySearchAddress(searchHistoryID, addressID int) error {
	_, err := gr.db.Exec("INSERT INTO history_search_address (search_history_id, address_id) VALUES ($1, $2)", searchHistoryID, addressID)
	if err != nil {
		return err
	}
	return nil
}


type AddressData struct {
    Region string
    GeoLat string
    GeoLon string
}


func (gr *geoRepository) FindAddressByQueryAndHistory(query string) ([]AddressData, error) {
    var addresses []AddressData
    // выполните запрос и сканируйте результат в структуру AddressData
    err := gr.db.Select(&addresses, `
        SELECT a.geo_lat as GeoLat, a.geo_lon as GeoLon, a.region as Region
        FROM address a
        JOIN history_search_address hsa ON a.id = hsa.address_id
        JOIN search_history sh ON sh.id = hsa.search_history_id
        WHERE sh.query LIKE $1
    `, "%"+query+"%")
    if err != nil {
        return nil, err
    }
    return addresses, nil
}





func (gr *geoRepository) SearchInHistory(query string) (bool, error) {
    var exists bool
    // Используем оператор % для поиска похожих запросов в таблице search_history
    err := gr.db.QueryRow("SELECT EXISTS (SELECT query FROM search_history WHERE query % $1)", query).Scan(&exists)
    if err != nil {
        return false, err
    }
    return exists, nil
}
