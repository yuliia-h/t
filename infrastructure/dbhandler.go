package infrastructure

import (
	"TestTest/interfaces"
	"github.com/jmoiron/sqlx"

	_ "database/sql"

	_ "github.com/lib/pq"
)

const schema = `CREATE TABLE if not exists images (
id    integer PRIMARY KEY GENERATED BY DEFAULT AS IDENTITY,
link  text NOT NULL,
UNIQUE(image)
)`

const makeunique = `ALTER TABLE images
ADD CONSTRAINT image UNIQUE (image);`

type DbimageConnect struct {
	dbimage *sqlx.DB
}

func NewDbimageConnect(dbimage *sqlx.DB) *DbimageConnect {
	return &DbimageConnect{dbimage: dbimage}
}

func (r DbimageConnect) MakeUniqueImageQuery() error {
	_, err := r.dbimage.Exec(makeunique)
	return err
}

func (r DbimageConnect) HistoryAll() ([]interfaces.ImageDb, error) {
	var images []interfaces.ImageDb

	err := r.dbimage.Select(&images, "select * from images")

	usercaseImages := make([]interfaces.ImageDb, 0, len(images))
	for i := range images {
		usercaseImages = append(usercaseImages, interfaces.ImageDb{Link: images[i].Link})
	}
	return usercaseImages, err
}

func (r DbimageConnect) FindImageId(id int) interfaces.ImageDb {
	temp := interfaces.ImageDb{}
	return temp
}

func (r DbimageConnect) ChangeImageId(id int) {
}

func (r DbimageConnect) SaveImage(image interfaces.ImageDb) (interfaces.ImageDb, error) {
	//запись данных обычная
	//_, err := r.dbimage.Exec("INSERT INTO images (width, height, link) VALUES ($1, $2, $3)",
	//	image.Width, image.Height, image.Link)

	//запись данных с получение Id
	var id int
	err := r.dbimage.QueryRow("INSERT INTO images (width, height, link) VALUES ($1, $2, $3) returning id",
		image.Width, image.Height, image.Link).Scan(&id)
	image.Id = id

	return image, err
}
