package webhandler

import (
	"log"
	"net/http"
	"regexp"

	"github.com/domainr/whois"
	"github.com/gin-gonic/gin"
	"github.com/schererja/domainaskfor-go/internal/app/domainaskfor/types"
	"github.com/schererja/domainaskfor-go/internal/app/domainaskfor/util"
)

func ShowHandler(context *gin.Context) {
	AvailableRegex := "No match for"
	UnavailableRegex := "Domain Name:"
	domain := util.GetURLFromReqContext(context)

	request, err := whois.NewRequest(domain)
	if err != nil {
		log.Print(err)
	}
	result, err := whois.DefaultClient.Fetch(request)
	if err != nil {
		log.Print(err)
	}
	availableMatched, err := regexp.MatchString(AvailableRegex, string(result.Body))
	if err != nil {
		log.Print(err)
	}
	unavailableMatch, err := regexp.MatchString(UnavailableRegex, string(result.Body))
	if err != nil {
		log.Print(err)
	}

	if availableMatched {
		context.JSON(http.StatusOK, types.WhoIsResult{
			DomainName:  domain,
			IsAvailable: true,
		})
	} else if unavailableMatch {
		context.JSON(http.StatusOK, types.WhoIsResult{
			DomainName:  domain,
			IsAvailable: false,
		})
	} else {
		context.JSON(http.StatusBadRequest, types.WhoIsResult{
			DomainName:  domain,
			IsAvailable: false,
		})
	}

}
