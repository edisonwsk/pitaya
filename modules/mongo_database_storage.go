package modules

import (
	"context"
	"github.com/qiniu/qmgo"
	"github.com/topfreegames/pitaya/config"
)


type MongoDatabaseStorage struct {
	Base
	config *config.Config

	client *qmgo.Client
	ctx context.Context

	dbConnect string
	dbName string
	dbUser string
	dbPwd string
}

func NewMongoDatabaseStorage(conf *config.Config)  *MongoDatabaseStorage {
	ds := &MongoDatabaseStorage{
		config:conf,
	}
	ds.configure()
	return ds
}

func (c *MongoDatabaseStorage)configure(){
	c.dbConnect = c.config.GetString("pitaya.modules.databasestorage.mango.connect")
	c.dbName = c.config.GetString("pitaya.modules.databasestorage.mango.dbname")
	c.dbUser = c.config.GetString("pitaya.modules.databasestorage.mango.user")
	c.dbPwd = c.config.GetString("pitaya.modules.databasestorage.mango.pwd")
}

// Init was called to initialize the component.
func (c *MongoDatabaseStorage) Init() error {
	c.ctx= context.Background()
	cfg :=&qmgo.Config{Uri: c.dbConnect}
	if len(c.dbUser)>0&&len(c.dbPwd)>0 {
		cfg.Auth = &qmgo.Credential{
			Username: c.dbUser,
			Password: c.dbPwd,
		}
	}
	client, err := qmgo.NewClient(c.ctx, cfg)
	if err!=nil {
		return  err
	}
	c.client =client
	return nil
}

// AfterInit was called after the component is initialized.
func (c *MongoDatabaseStorage) AfterInit() {

}

// BeforeShutdown was called before the component to shutdown.
func (c *MongoDatabaseStorage) BeforeShutdown() {

}

// Shutdown was called to shutdown the component.
func (c *MongoDatabaseStorage) Shutdown() error {
	return c.client.Close(c.ctx)
}


func (c *MongoDatabaseStorage) GetDB() *qmgo.Database{
	return c.client.Database(c.dbName)
}



