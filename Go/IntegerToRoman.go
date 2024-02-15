type Roman struct {
    Key string
    Value int
    BeforeKey string
    BeforeValue int
}

func intToRoman(num int) string {

    response := ""    
    romanValues := []Roman{
        {Key: "M", Value: 1000, BeforeKey: "C", BeforeValue: 900},
        {Key: "D", Value: 500, BeforeKey: "C", BeforeValue: 400},
        {Key: "C", Value: 100, BeforeKey: "X", BeforeValue: 90},
        {Key: "L", Value: 50, BeforeKey: "X", BeforeValue: 40},
        {Key: "X", Value: 10, BeforeKey: "I", BeforeValue: 9},
        {Key: "V", Value: 5, BeforeKey: "I", BeforeValue: 4},
        {Key: "I", Value: 1}}

    for _, v := range romanValues {

        for num >= v.Value {
            response += v.Key
            num -= v.Value
        }

        if num >= v.BeforeValue && v.BeforeValue > 0 {
            response += v.BeforeKey + v.Key
            num -= v.BeforeValue
        }

    }

    return response

}