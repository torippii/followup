package util

import (
	"encoding/csv"
	"encoding/json"
	"hrmos-bulk-register/model"
	"io"
	"log"
	"os"
	"strings"
)

const (
	ATTENDANCE_CSV = "./attendance.csv"
	SETTING_JSON = "./setting.json"
	COMMENT = "#"
)

func LoadCSV() [] model.Attendance{
	 f, err := os.Open(ATTENDANCE_CSV)
	 if err != nil {
		log.Fatalln(err)
	 }
	 defer f.Close()

	 var attendanceList [] model.Attendance
	 r := csv.NewReader(f)
	 for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		// csvの中で先頭行に'#'がある場合はコメントアウト行とみなす
		if strings.Contains(record[0], COMMENT) {
			continue
		}

		attendance := model.NewAttendance(record[0], record[1], record[2], record[3], record[4])
		attendanceList = append(attendanceList, *attendance)
	 }
	 return attendanceList
}

func LoadSettingJson() model.User{
	f, err := os.Open(SETTING_JSON)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	var user model.User
	if err := json.NewDecoder(f).Decode(&user); err != nil {
		log.Fatalln(err)
	}

	return user
}