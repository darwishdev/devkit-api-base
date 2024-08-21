package convertor

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToPgType(str string) pgtype.Text {

	return pgtype.Text{String: str, Valid: str != ""}
}

func ToPgTimeString(pgTime pgtype.Time) string {
	duration := time.Duration(pgTime.Microseconds) * time.Microsecond

	// Convert duration to time.Time
	timeValue := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC).Add(duration)

	// Format time as "15:04" (24-hour clock format)
	return timeValue.Format("15:04")
}

func ToPgDate(strDate string) pgtype.Date {
	parsedTime, _ := time.Parse("2006-01-02", strDate)

	year, month, day := parsedTime.Date()

	// Create pgtype.Date
	pgDate := pgtype.Date{
		Time:  time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
		Valid: true,
	}

	return pgDate
}
func ToPgTimeStamp(dateString string) pgtype.Timestamp {
	// Step 1: Parse the string into a time.Time
	fmt.Println("Error parsing time:")
	parsedTime, err := time.Parse("1/2/2006, 3:04:05 PM", dateString)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return pgtype.Timestamp{Valid: false}
	}

	// Step 2: Create a pgtype.Timestamp using the parsed time
	pgTimestamp := pgtype.Timestamp{Time: parsedTime, Valid: true}
	return pgTimestamp
}
func ToPgTime(strTime string) pgtype.Time {
	parsedTime, err := time.Parse("15:04", strTime)
	if err != nil {
		log.Debug().Interface("errerrerrerrerrerr", err).Msg("Error parsing time:")
		return pgtype.Time{Valid: false}
	}

	// Extract hour, minute, second, and microsecond components
	hour, minute, second := parsedTime.Clock()
	microseconds := parsedTime.Nanosecond() / 1000 // Convert nanoseconds to microseconds

	// Create pgtype.Time
	pgTime := pgtype.Time{
		Microseconds: int64(hour*3600+minute*60+second)*1e6 + int64(microseconds),
		Valid:        true,
	}

	return pgTime
}
func ToPgTypeBool(value bool) pgtype.Bool {
	return pgtype.Bool{Bool: value, Valid: true}
}
func ToPgTypeInt(value int32) pgtype.Int4 {
	return pgtype.Int4{Int32: value, Valid: true}
}
func ToPgTypeFloat(value float32) pgtype.Float4 {
	return pgtype.Float4{Float32: value, Valid: true}
}
func ToPgTypeID(value int32) pgtype.Int4 {
	return pgtype.Int4{Int32: value, Valid: value > 0}
}
func ToPgTypeUInt(value int32) pgtype.Int4 {
	return pgtype.Int4{Int32: value, Valid: value > -1}
}
func ToTimeStamp(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}
func ToSelectInput(value int32, label string, icon string, note string) *abcv1.SelectInputOption {
	return &abcv1.SelectInputOption{
		Value: value,
		Label: label,
		Icon:  icon,
		Note:  note,
	}
}

func ToSelectInputWithGroup(value int32, label string, icon string, note string, groupName string, groupIcon string, resultMap map[string]*abcv1.SelectInputOptionWithGroup) bool {
	log.Debug().Interface("grouname", groupName).Msg("groups")
	optionItem := ToSelectInput(value, label, icon, note)
	_, ok := resultMap[groupName]
	if !ok {
		resultMap[groupName] = &abcv1.SelectInputOptionWithGroup{
			GroupName: groupName,
			GroupIcon: groupIcon,
			Items: []*abcv1.SelectInputOption{
				optionItem,
			},
		}

		return true
	}

	resultMap[groupName].Items = append(resultMap[groupName].Items, optionItem)
	return false

}

func SnakeToPascal(input string) string {
	words := strings.Split(input, "_")
	for i := range words {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}

func SetField(obj interface{}, name string, value interface{}) error {
	// Get the struct value from the interface
	structValue := reflect.ValueOf(obj).Elem()

	// Get the field value by name
	fieldValue := structValue.FieldByName(name)

	// Check if the field exists
	if !fieldValue.IsValid() {
		return fmt.Errorf("no such field: %s in obj", name)
	}

	// Get the type of the field
	fieldType := fieldValue.Type()

	// Get the value to be set as a reflect.Value
	newValue := reflect.ValueOf(value)

	// Check if the type of the value to be set matches the type of the field
	if !newValue.Type().AssignableTo(fieldType) {
		return fmt.Errorf("value type %v is not assignable to field type %v", newValue.Type(), fieldType)
	}

	// Set the field value to the new value
	fieldValue.Set(newValue)

	return nil
}
