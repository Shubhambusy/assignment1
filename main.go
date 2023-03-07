// You have to iterate input and print the type and value of the entity.
// If it is a data structure then go inside it and do the same.
// You have to keep on going nested until you reach data types.
// (Do this using reflect concept)
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

func checkvaluetype (value any) (error) {
    // printing data type
    fmt.Println("Data Type: ", reflect.TypeOf(value).Kind())

    switch reflect.TypeOf(value).Kind() {
    case reflect.Slice:
        arr, e := value.([]any)
        if (e == false) {
            return errors.New("Invalid JSON array")
        }
        for _, c := range arr {
            e := checkvaluetype(c)
            if (e != nil) {
                return e
            }
        }
    case reflect.Map:

        // Here, we can directly convert value to map[string]interface{}
        // but that way we will not be able to validate the json
        innerjsondata, e := json.Marshal(value)
        if e != nil {return e}
        readJSON(innerjsondata)
    default:
        // printing values 
        fmt.Println("value: ", value)
    }
    fmt.Println(".................................................")

    return nil
}

func readJSON(jsondata []byte) (error) {

    // a map container to decode the JSON structure into
    data := make(map[string]interface{})

    // converting JSON to map[string]interface{}
    e := json.Unmarshal(jsondata, &data)
    if e != nil {
        return e
    }

    for key, value := range data {  // key, value
        fmt.Println("key", key)
        
        e = checkvaluetype(value)
        if e != nil {
            return e
        }
    }
    return nil
}

func main() {
    data := []byte(`{"name" : "Tolexo Online Pvt. Ltd","age_in_years" : 8.5,"origin" : "Noida","head_office" : "Noida, Uttar Pradesh","address" : [{"street" : "91 Springboard","landmark" : "Axis Bank","city" : "Noida","pincode" : 201301,"state" : "Uttar Pradesh"},{"street" : "91 Springboard","landmark" : "Axis Bank","city" : "Noida","pincode" : 201301,"state" : "Uttar Pradesh"}],"sponsers" : {"name" : "One"},"revenue" : "19.8 million$","no_of_employee" : 630,"str_text" : ["one","two"],"int_text" : [1,3,4]}`)
    var e error
    if e = readJSON(data); e != nil {
        fmt.Println("The given JSON string was not the valid json", e)
    }
}
