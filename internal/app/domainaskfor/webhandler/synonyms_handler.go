package webhandler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/schererja/domainaskfor-go/internal/app/domainaskfor/types"
	"github.com/schererja/domainaskfor-go/internal/app/domainaskfor/util"
)

func ShowSynonymsHandler(context *gin.Context) {
	var synonymsFound []string

	word := util.GetWordFromReqContext(context)
	cachedSynonyms := util.GetCachedSynonyms(word)
	log.Print(len(cachedSynonyms))
	if len(cachedSynonyms) <= 1 {
		dictionaryAPIUrl := strings.Replace("https://dictionaryapi.com/api/v3/references/thesaurus/json/*?key="+os.Getenv("dictionarykey"), "*", word, -1)
		resp, err := http.Get(dictionaryAPIUrl)
		if err != nil {
			log.Print(err)
		}
		defer resp.Body.Close()
		var webster []types.WebsterResult
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Print(err)
		}
		err = json.Unmarshal(bytes, &webster)
		if err != nil {
			log.Print(err)
		}
		for i := 0; i < len(webster); i++ {
			syns := webster[i].Meta.Syns
			for k := 0; k < len(syns); k++ {
				for j := 0; j < len(syns[k]); j++ {
					spaceRemoved := strings.ReplaceAll(syns[k][j], " ", "")
					synonymsFound = append(synonymsFound, spaceRemoved)
				}

			}
		}

		util.SetCachedSynonyms(word, synonymsFound)
	} else {
		synonymsFound = cachedSynonyms
	}
	context.JSON(http.StatusOK, struct {
		Synonyms []string
	}{
		Synonyms: synonymsFound,
	})

}
