package helpers

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"github.com/mehrdadnekopour/go-tools/mypes"

	"github.com/fatih/structs"
)

// Function pluck is used to retrieve an array of subset of fields(branch) present in original structure(plant).
// Input : 'plant' is the source from which a branch needs to be plucked. An array of structure is expected.
//         'branch' is the output structure(subset). This should not be an array as this will be used just to
//              form the output structure.
// Output : []map[string]interface{} - An array of map is returned. See example for more details.

// Pluck ...
func Pluck(plant interface{}, branch interface{}) ([]map[string]interface{}, error) {

	flag := 0
	// Read the value from interface{}
	srcExtract := reflect.ValueOf(plant)

	// For branch, only format is needed(and the input is not an array)
	// The value is extracted and converted to map.
	destExtract := reflect.ValueOf(branch).Interface()
	destValMap := structs.Map(destExtract)

	// The result map[string]interface{} to be returned
	branchExtract := make([]map[string]interface{}, srcExtract.Len())

	// Retrieve the source elements one by one and copy to dest
	for i := 0; i < srcExtract.Len(); i++ {
		indexVal := srcExtract.Index(i).Interface()
		indexValMap := structs.Map(indexVal)
		// Create a temp variable to hold the trimmed values
		destTempMap := make(map[string]interface{})
		for key := range destValMap {
			if value, present := indexValMap[key]; present {
				destTempMap[key] = value
				flag = 1
			}
		}
		// append the temp var into output
		branchExtract[i] = destTempMap
	}
	// This is to make sure that at least one value got extracted from plant to branch
	if flag == 0 {
		err := errors.New("Source Destination Type Mismatch")
		return nil, err
	}
	return branchExtract, nil
}

// Function pluckElement is used to retrieve an array of just one field(destKeyName) present in original structure(plant).
// Input : 'plant' is the source from which an element needs to be plucked. An array of structure is expected.
//         'destKeyName' is the output element key name. This should not be an array as this will be used just to
//              form the output structure.
// Output : []interface{} - An array is returned. Type assertion can be used to derive an array of required type.
// See example for more details.

// PluckElement ...
func PluckElement(plant interface{}, destKeyName string) ([]interface{}, error) {

	flag := 0
	// Read the value from interface{}
	srcExtract := reflect.ValueOf(plant)

	// The result map[string]interface{} to be returned
	var elementExtract []interface{}

	// Retrieve the source elements one by one and copy to dest
	for i := 0; i < srcExtract.Len(); i++ {
		indexVal := srcExtract.Index(i).Interface()
		indexValMap := structs.Map(indexVal)
		if value, present := indexValMap[destKeyName]; present {
			elementExtract = append(elementExtract, value)
			flag = 1
		}
	}

	// This is to make sure that at least one value got extracted from plant to branch
	if flag == 0 {
		err := errors.New("Source Destination Type Mismatch")
		return nil, err
	}
	return elementExtract, nil
}

// Cast ...
func Cast(src interface{}, dst interface{}) {
	dstValue := reflect.ValueOf(dst)
	dstElem := dstValue.Elem()
	countDstElems := dstElem.NumField()
	dstType := dstElem.Type()

	for i := 0; i < countDstElems; i++ {
		dstField := dstElem.Field(i)
		dstFieldName := dstType.Field(i).Name
		// dstFieldType := dstField.Type()

		srcValue := reflect.ValueOf(src)
		srcElem := srcValue.Elem()
		countSrcElems := srcElem.NumField()
		srcType := srcElem.Type()

		if i == 0 {
			valID := reflect.Indirect(srcValue).Field(i)
			dstField.Set(valID)
			continue
		}

		for j := 1; j < countSrcElems; j++ {

			// srcField := srcElem.Field(j)
			srcFieldName := srcType.Field(j).Name
			// srcFieldType := srcField.Type()

			if dstFieldName == srcFieldName { // && dstFieldType == srcFieldType {
				// FOUND
				val := reflect.Indirect(srcValue).Field(j)
				// f := reflect.Indirect(r).FieldByName("Mobile")
				dstField.Set(val)
				break
			}
		}
	}

	// s := reflect.ValueOf(src).Elem()

	// fmt.Println(s.NumField())
	// typeOfR := s.Type()

	// for i := 0; i < s.NumField(); i++ {
	// 	f := s.Field(i)
	// 	fmt.Printf("%d: %s %s = %v\n", i, typeOfR.Field(i).Name, f.Type(), f.Interface())
	// }
}

// CastToInt ...
func CastToInt(iface interface{}) (int, error) {
	// iaface = indirect(iface)

	if iface == nil {
		return 0, nil
	}

	var err error
	switch s := iface.(type) {
	case string:
		return strconv.Atoi(iface.(string))
	case int:
		return s, err
	case int32:
		return int(s), err
	case int16:
		return int(s), err
	case int8:
		return int(s), err
	case uint:
		return int(s), err
	case uint64:
		return int(s), err
	case uint32:
		return int(s), err
	case uint16:
		return int(s), err
	case uint8:
		return int(s), err
	case float64:
		return int(s), err
	case float32:
		return int(s), err
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int", iface, iface)
	}
}

// CastToDouble ...
func CastToDouble(iface interface{}) (float64, error) {
	// iaface = indirect(iface)

	if iface == nil {
		return 0, nil
	}

	// var err error
	switch iface.(type) {
	case string:
		{
			d, err := strconv.ParseFloat(iface.(string), 64)
			return d, err
		}
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int", iface, iface)
	}
}

// GetUintFromString ...
func GetUintFromString(s string) (uint, error) {
	idP, err := strconv.Atoi(s)
	uID := uint(idP)

	return uID, err
}

// GetLimitOffsetFromMSON ...
func GetLimitOffsetFromMSON(params mypes.MSON) (limit, offset int, e error) {
	limit, e = CastToInt(params["limit"])
	if e != nil {
		return
	}

	offset, e = CastToInt(params["offset"])
	if e != nil {
		return
	}

	if limit > 50 {
		limit = 50
	} else if limit < 0 {
		limit = 50
	}

	if offset < 0 {
		offset = 0
	}

	return
}

//CreateDirIfNotExist ...
func CreateDirIfNotExist(assetsPath, dir string) (os.FileInfo, string, error) {
	// rootDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// d := rootDir + dir

	d := assetsPath + dir
	fmt.Println(d)
	fInfo, err := os.Stat(d)

	if err == nil {
		return fInfo, d, err
	}

	if os.IsNotExist(err) {
		err = os.MkdirAll(d, 0755)
		if err != nil {
			return fInfo, "", err
		}
	}
	return fInfo, d, err
}
