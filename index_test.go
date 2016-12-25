package prettybytes

import "testing"
import "github.com/stretchr/testify/assert"

func TestPretty(t *testing.T) {
  assert := assert.New(t)
  var pretty string

  pretty, _ = Pretty(0)
  assert.Equal("0 B", pretty, "they should be equal")
  pretty, _ = Pretty(0.4)
  assert.Equal("0.4 B", pretty, "they should be equal")
  pretty, _ = Pretty(0.7)
  assert.Equal("0.7 B", pretty, "they should be equal")
  pretty, _ = Pretty(10)
  assert.Equal("10.00 B", pretty, "they should be equal")
  pretty, _ = Pretty(10.1)
  assert.Equal("10.10 B", pretty, "they should be equal")
  pretty, _ = Pretty(999)
  assert.Equal("999.00 B", pretty, "they should be equal")
  pretty, _ = Pretty(1001)
  assert.Equal("1.00 kB", pretty, "they should be equal")
  pretty, _ = Pretty(1e16)
  assert.Equal("10.00 PB", pretty, "they should be equal")
  pretty, _ = Pretty(1e30)
  assert.Equal("1000000.00 YB", pretty, "they should be equal")

  pretty, _ = Pretty(-0.4)
  assert.Equal("-0.4 B", pretty, "they should be equal")
  pretty, _ = Pretty(-0.7)
  assert.Equal("-0.7 B", pretty, "they should be equal")
  pretty, _ = Pretty(-10.1)
  assert.Equal("-10.10 B", pretty, "they should be equal")
  pretty, _ = Pretty(-999)
  assert.Equal("-999.00 B", pretty, "they should be equal")
  pretty, _ = Pretty(-1001)
  assert.Equal("-1.00 kB", pretty, "they should be equal")
}