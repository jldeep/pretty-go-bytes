package prettybytes

import "math"
import "strconv"
import "errors"

func Pretty (number float64) (string, error) {
  UNITS := [9]string{"B", "kB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
  var value string
  val := math.IsNaN(number)
  if (val) {
      return "", errors.New("Number Is NaN")
  }  

  neg := number < 0

  if neg {
      number = -number
  }

  if number < 1 {
    if neg {
      value += "-"
    }
    value += strconv.FormatFloat(number, 'f', -1, 64) + " B"
    return value, nil
  }

  exponent := math.Min(math.Floor( math.Log(number) / math.Log(1000) ), float64(len(UNITS) - 1))
  numStr := (number / math.Pow(1000, exponent))
  unit := UNITS[int(exponent)]
  if neg {
    value += "-"
  }
  value += strconv.FormatFloat(numStr, 'f', 2, 64) + " " + unit
  return value, nil
}