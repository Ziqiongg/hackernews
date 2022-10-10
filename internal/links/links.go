package links

import (
	"errors"
	// "fmt"
	// "log"
	// "strings"
	"gorm.io/gorm"
	//"gorm.io/gorm"
	//"github.com/ziqiongg/hackernews/server.go"
)

type Link struct {
	ID      int64  `gorm:"id"`
	Title   string `gorm:"title"`
	Address string `gorm:"address"`
	UserID  int64  `gorm:"user_id"`
}

// type LinkModel struct{
// 	gorm.Model
// 	*Link
// }

func (link Link) SaveLink(Db *gorm.DB) int64 {
	linkexample := &Link{ID: 1,
		Title:   link.Title,
		Address: link.Address}

	err := Db.Create(linkexample).Error
	if err != nil {
		errors.New("cant create new links")
	}

	return linkexample.ID

}

// func GetAll() []Link {
// 	stmt, err := db.Db.Prepare("select id, title, address from Links")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer stmt.Close()
// 	rows, err := stmt.Query()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()
// 	var links []Link
// 	for rows.Next() {
// 		var link Link
// 		err := rows.Scan(&link.ID, &link.Title, &link.Address)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		links = append(links, link)
// 	}
// 	if err = rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 	return links
// }
