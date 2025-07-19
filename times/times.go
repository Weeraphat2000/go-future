package times

import (
	"fmt"
	"time"
)

// GetCurrentTimeAndFormat returns the current time as a formatted string
func GetCurrentTimeAndFormat() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GetCurrentTime returns the current time as a time.Time object
func GetCurrentTime() time.Time {
	return time.Now()
}

func PrintCurrentTime() {
	fmt.Println("Current time is:", GetCurrentTimeAndFormat())
	fmt.Println("time.Now()", GetCurrentTime())
}

func FormatTimeExample() {
	fmt.Println(" ")
	fmt.Println("Format Time Example")

	t := time.Date(2023, time.November, 20, 15, 30, 0, 0, time.UTC)
	fmt.Println("Specific time:", t)

	// TODO: set format ต้องเป็นแบบนี้เท่านั้น
	// "2006-01-02 15:04:05" เป็น format ที่กำหนดไว้ใน Go
	format := t.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted specific time:", format)

	// TODO: set format ต้องเป็นแบบนี้เท่านั้น
	// "02 Jan 2006 15:04:05" เป็น format ที่กำหนดไว้ใน Go
	format2 := t.Format("02 Jan 2006 15:04:05")
	fmt.Println("Another formatted specific time:", format2)
}

func ParseTimeString() {
	fmt.Println(" ")
	fmt.Println("Parse Time String Example")

	// TODO: set format ต้องเป็นแบบนี้เท่านั้น
	// "2006-01-02T15:04:05" เป็น format ที่กำหนดไว้ใน Go
	format3 := "2006-01-02T15:04:05"
	// TODO: เช็คว่า time string ตรงกับ format นี้หรือไม่
	t2, err := time.Parse(format3, "2023-11-11T00:00:00")
	if err != nil {
		// format ผิด
		fmt.Println("Error parsing time22222:", err)
	} else {
		fmt.Println("Parsed time22222:", t2)
	}

	format4 := "2006-01-02T15:04:05"
	// TODO: เช็คว่า time string ตรงกับ format นี้หรือไม่
	t3, err4 := time.Parse(format4, "02 Jan 2006 15:04:05")
	if err4 != nil {
		// format ผิด
		fmt.Println("Error parsing time3333:", err4)
	} else {
		fmt.Println("Parsed time3333:", t3)
	}

	format5 := "02 Jan 2006"
	// TODO: เช็คว่า time string ตรงกับ format นี้หรือไม่
	t4, err5 := time.Parse(format5, "20 Nov 2023")
	if err5 != nil {
		// format ผิด
		fmt.Println("Error parsing time4444:", err5)
	} else {
		fmt.Println("Parsed time4444:", t4)
	}
}

func AddTime() {
	fmt.Println(" ")
	fmt.Println("Add Time Example")

	now := time.Now()
	fmt.Println("Current time:", now)

	// TODO: add 1 hour
	oneHourLater := now.Add(time.Hour)
	fmt.Println("One hour later:", oneHourLater)

	// TODO: add 30 minutes
	thirtyMinutesLater := now.Add(30 * time.Minute)
	fmt.Println("Thirty minutes later:", thirtyMinutesLater)

	// TODO: add 15 seconds
	fifteenSecondsLater := now.Add(15 * time.Second)
	fmt.Println("Fifteen seconds later:", fifteenSecondsLater)
}

func SubtractTime() {

	fmt.Println(" ")
	fmt.Println("Subtract Time Example")

	now := time.Now()
	fmt.Println("Current time:", now)

	// TODO: subtract 1 hour
	oneHourEarlier := now.Add(-time.Hour)
	fmt.Println("One hour earlier:", oneHourEarlier)

	// TODO: subtract 30 minutes
	thirtyMinutesEarlier := now.Add(-30 * time.Minute)
	fmt.Println("Thirty minutes earlier:", thirtyMinutesEarlier)

	// TODO: subtract 15 seconds
	fifteenSecondsEarlier := now.Add(-15 * time.Second)
	fmt.Println("Fifteen seconds earlier:", fifteenSecondsEarlier)
}

func CompareTimes() {

	fmt.Println(" ")
	fmt.Println("Compare Times Example")
	t := time.Date(2023, time.October, 1, 12, 30, 0, 0, time.UTC)
	now := time.Now()
	fmt.Println("Current time:", now)
	fmt.Println("Specific time:", t)

	fmt.Println("Is current time after specific time?", now.After(t))           // true
	fmt.Println("Is current time before specific time?", now.Before(t))         // false
	fmt.Println("Is current time equal to specific time?", now.Equal(t))        // false
	fmt.Println("Duration between current time and specific time:", now.Sub(t)) // Duration between current time and specific time

	t2 := time.Date(2023, time.October, 1, 12, 30, 0, 0, time.UTC)
	t3 := time.Date(2023, time.October, 1, 12, 30, 0, 0, time.UTC)
	fmt.Println("Is t2 after t3?", t2.After(t3))           // false
	fmt.Println("Is t2 before t3?", t2.Before(t3))         // false
	fmt.Println("Is t2 equal to t3?", t2.Equal(t3))        // true
	fmt.Println("Duration between t2 and t3:", t2.Sub(t3)) // 0s
	fmt.Println("Duration between t3 and t2:", t3.Sub(t2)) // 0s
}

func DurationExample() {
	start := time.Now()
	// TODO: sleep for 2 seconds to simulate work
	time.Sleep(2 * time.Second)
	end := time.Now()

	duration := end.Sub(start)
	fmt.Println("Duration of work:", duration)
}

func TimeZoneExample() {
	fmt.Println(" ")
	fmt.Println("Time Zone Example")
	if loc, err := time.LoadLocation("Asia/Bangkok"); err != nil {
		fmt.Println("Error loading location:", err)
	} else {
		t := time.Date(2023, time.October, 1, 12, 30, 0, 0, loc)
		fmt.Println("Specific time in Asia/Bangkok:", t)
		fmt.Println("Current time in Asia/Bangkok:", time.Now().In(loc))
	}

	fmt.Println("Current time in UTC:", time.Now().UTC())
	if loc, err := time.LoadLocation("America/New_York"); err != nil {
		fmt.Println("Error loading location:", err)
	} else {
		t := time.Date(2023, time.October, 1, 12, 30, 0, 0, loc)
		fmt.Println("Specific time in America/New_York:", t)
		fmt.Println("Current time in America/New_York:", time.Now().In(loc))
	}
}

func TimeExample() {

	fmt.Println(" ")
	fmt.Println("Time Example")
	now := time.Now()
	fmt.Println("Current time:", now)

	// TODO: create time
	t := time.Date(2023, time.October, 1, 12, 30, 0, 0, time.UTC)
	fmt.Println("Specific time:", t)

	FormatTimeExample()

	ParseTimeString()

	AddTime()

	SubtractTime()

	CompareTimes()

	DurationExample()

	TimeZoneExample()
}
