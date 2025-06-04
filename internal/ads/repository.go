package ads

import (
	"database/sql"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetAllAds() ([]Ad, error) {
	rows, err := r.db.Query("SELECT id, image_url, target_url FROM ads")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ads []Ad
	for rows.Next() {
		var a Ad
		rows.Scan(&a.ID, &a.ImageURL, &a.TargetURL)
		ads = append(ads, a)
	}
	return ads, nil
}

func (r *Repository) InsertClick(c Click) error {
	_, err := r.db.Exec("INSERT INTO clicks(ad_id, timestamp, ip, playback_time) VALUES ($1, $2, $3, $4)",
		c.AdID, time.Now(), c.IP, c.PlaybackTime)
	return err
}

func (r *Repository) GetClickCounts() (map[int]int, error) {
	rows, err := r.db.Query("SELECT ad_id, COUNT(*) FROM clicks GROUP BY ad_id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	counts := make(map[int]int)
	for rows.Next() {
		var adID, count int
		rows.Scan(&adID, &count)
		counts[adID] = count
	}
	return counts, nil
}
