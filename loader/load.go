package loader

import (
	"github.com/jinzhu/gorm"
	"github.com/neiln3121/music-service/models"
)

func LoadData(db *gorm.DB) {
	db.Create(&models.Artist{Name: "Prince & the Revolution", Albums: []*models.Album{
		{Title: "Purple Rain", Year: 1984,
			Tracks: []*models.Track{{Title: "Let's Go Crazy"}, {Title: "Take Me With U"}, {Title: "The Beautiful Ones"}, {Title: "Computer Blue"},
				{Title: "Darling Nikki"}, {Title: "When Doves Cry"}, {Title: "I Would Die 4 U"}, {Title: "Baby, I'm a Star"}, {Title: "Purple Rain"}}},
		{Title: "Parade", Year: 1986,
			Tracks: []*models.Track{{Title: "Christopher Tracy's Parade"}, {Title: "New Position"}, {Title: "I Wonder U"}, {Title: "Under the Cherry Moon"}, {Title: "Girls & Boys"},
				{Title: "Venus de Milo"}, {Title: "Mountains"}, {Title: "Do U Lie?"}, {Title: "Kiss"}, {Title: "Anotherloverholenyohead"}, {Title: "Sometimes It Snows in April"}}},
		{Title: "Around the World in a Day", Year: 1985, Tracks: []*models.Track{{Title: "Around the World in a Day"}, {Title: "Paisley Park"}, {Title: "Condition of the Heart"}, {Title: "Raspberry Beret"},
			{Title: "Tamborine"}, {Title: "Pop Life"}, {Title: "The Ladder"}, {Title: "Temptation"}}},
	}})

	db.Create(&models.Artist{Name: "INXS", Albums: []*models.Album{
		{Title: "Kick", Year: 1987,
			Tracks: []*models.Track{{Title: "Guns in the Sky"}, {Title: "New Sensation"}, {Title: "Devil Inside"}, {Title: "Need You Tonight"}, {Title: "Mediate"}, {Title: "The Loved One"},
				{Title: "Wild Life"}, {Title: "Never Tear Us Apart"}, {Title: "Mystify"}, {Title: "Kick"}, {Title: "Calling All Nations"}, {Title: "Tiny Daggers"}}}}})

	db.Create(&models.Artist{Name: "U2", Albums: []*models.Album{
		{Title: "The Joshua Tree", Year: 1987,
			Tracks: []*models.Track{{Title: "Where the Streets Have No Name"}, {Title: "I Still Haven't Found What I'm Looking For"}, {Title: "With or Without You"}, {Title: "Bullet the Blue Sky"},
				{Title: "Running to Stand Still"}, {Title: "Red Hill Mining Town"}, {Title: "In God's Country"}, {Title: "Trip Through Your Wires"}, {Title: "One Tree Hill"}, {Title: "Exit"},
				{Title: "Mothers of the Disappeared"}}},

		{Title: "Achtung Baby", Year: 1991,
			Tracks: []*models.Track{{Title: "Zoo Station"}, {Title: "Even Better Than the Real Thing"}, {Title: "One"}, {Title: "Until the End of the World"}, {Title: "Who's Gonna Ride Your Wild Horses"},
				{Title: "So Cruel"}, {Title: "The Fly"}, {Title: "Mysterious Ways"}, {Title: "Tryin' to Throw Your Arms Around the World"}, {Title: "Ultraviolet (Light My Way)"}, {Title: "Acrobat"},
				{Title: "Love Is Blindness"}}},
	}})
}
