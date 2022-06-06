package restclient

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_main(t *testing.T) {
	err, statusCode := main()

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, statusCode)

}
