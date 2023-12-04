package queries

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type Users struct {
	username     string
	date_created string
}

type Followers struct {
	username      string
	followerCount int
}

type Likes struct {
	username string
	postID   string
	likes    int
}

func BlueButton() (string, []Users) {
	// start := time.Now()
	connectionString := "user=matthewbruner host=localhost port=5432 dbname=python-vs-go sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	var records []Users
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("DB Connection: %v\n", time.Since(start))

	// start = time.Now()
	query := "SELECT username, date_created FROM users;"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// fmt.Printf("Query DB: %v\n", time.Since(start))

	swapHTML := `<table class='table-auto w-full border border-collapse text-center'>
	<thead>
	<tr>
	<th class='px-4 py-2 border border-gray-600 bg-slate-400'>Username</th>
	<th class='px-4 py-2 border border-gray-600 bg-slate-400'>Account Creation Date</th>
	</tr>
	</thead>
	<tbody>`

	// start = time.Now()
	var (
		column1 string
		column2 time.Time
	)
	for rows.Next() {
		if err := rows.Scan(&column1, &column2); err != nil {
			log.Fatal(err)
		}

		records = append(records, Users{
			username:     column1,
			date_created: column2.Format("2006-01-02"),
		})
	}
	// fmt.Printf("Unpack all Data: %v\n", time.Since(start))

	for i := 0; i < 5; i++ {
		swapHTML = swapHTML + fmt.Sprintf(`
		<tr>
		<td class='px-4 py-2 border border-gray-600'>%s</td>
		<td class='px-4 py-2 border border-gray-600'>%s</td>
		</tr>`,
			records[i].username, records[i].date_created)
	}
	swapHTML = swapHTML + "</tbody></table>"

	return swapHTML, records
}

func RedButton() (string, []Followers) {
	connectionString := "user=matthewbruner host=localhost port=5432 dbname=python-vs-go sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	var records []Followers

	query := `
		SELECT u.username, count(u.username) AS follower_count
			FROM users AS u
			INNER JOIN followers as f
			ON u.id = f.user_id
			GROUP BY u.username
			ORDER BY follower_count DESC;
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	swapHTML := `<table class='table-auto w-full border	border-gray-600 text-center'>
	<thead>
	<tr>
	<th class='px-4 py-2 border-separate border-gray-600 bg-slate-400'>Username</th>
	<th class='px-4 py-2 border-separate border-gray-600 bg-slate-400'>Follower Count</th>
	</tr>
	</thead>
	<tbody>`

	var (
		column1 string
		column2 int
	)
	for rows.Next() {
		if err := rows.Scan(&column1, &column2); err != nil {
			log.Fatal(err)
		}

		records = append(records, Followers{
			username:      column1,
			followerCount: column2,
		})
	}

	for i := 0; i < 5; i++ {
		swapHTML = swapHTML + fmt.Sprintf(`
		<tr>
		<td class='px-4 py-2 border border-gray-600'>%s</td>
		<td class='px-4 py-2 border border-gray-600'>%v</td>
		</tr>`,
			records[i].username, records[i].followerCount)
	}
	swapHTML = swapHTML + "</tbody></table>"

	return swapHTML, records
}

func PurpleButton() (string, []Likes) {
	connectionString := "user=matthewbruner host=localhost port=5432 dbname=python-vs-go sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	var records []Likes

	query := `
		SELECT u.username, p.id, count(l.post_id) AS total_likes
			FROM likes AS l
			LEFT JOIN posts AS p
			ON l.post_id = p.id
			LEFT JOIN users AS u
			ON u.id = p.user_id
			GROUP BY p.id, u.username
			ORDER BY total_likes DESC;
	`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	swapHTML := `<table class='table-auto w-full border	border-gray-600 text-center'>
	<thead>
	<tr>
	<th class='px-4 py-2 border-separate border-gray-600 bg-slate-400'>Username</th>
	<th class='px-4 py-2 border-separate border-gray-600 bg-slate-400'>Post ID</th>
	<th class='px-4 py-2 border-separate border-gray-600 bg-slate-400'>Total Likes</th>
	</tr>
	</thead>
	<tbody>`

	var (
		column1 string
		column2 string
		column3 int
	)
	for rows.Next() {
		if err := rows.Scan(&column1, &column2, &column3); err != nil {
			log.Fatal(err)
		}

		records = append(records, Likes{
			username: column1,
			postID:   column2,
			likes:    column3,
		})
	}

	for i := 0; i < 5; i++ {
		swapHTML = swapHTML + fmt.Sprintf(`
		<tr>
		<td class='px-4 py-2 border border-gray-600'>%s</td>
		<td class='px-4 py-2 border border-gray-600'>%s</td>
		<td class='px-4 py-2 border border-gray-600'>%v</td>
		</tr>`,
			records[i].username, records[i].postID, records[i].likes)
	}
	swapHTML = swapHTML + "</tbody></table>"

	return swapHTML, records
}

func ClearButton() string {
	return `
	<table class="table-auto w-full border border-collapse text-center" id="target">
		<thead>
			<tr>
				<th class="px-4 py-2 border border-gray-600">Username</th>
				<th class="px-4 py-2 border border-gray-600">Account Creation Date</th>
			</tr>
		</thead>
	</table>
	`
}

func ExportButton(latencies []string) string {
	latencyValues := latencies
	file, err := os.Create("latencies/output.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range latencyValues {
		err := writer.Write([]string{value})
		if err != nil {
			log.Fatal(err)
		}
	}
	return "Export Latencies"
}
