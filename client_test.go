package odesli

import (
	"encoding/json"
	"log"
	"testing"
)

func TestGetLinksResponse_Unmarshall_IDIsString(t *testing.T) {
	data := `{
    "entityUniqueId": "ITUNES_ALBUM::1751932130",
    "userCountry": "US",
    "pageUrl": "https://album.link/us/i/1751932130",
    "entitiesByUniqueId":
    {
        "AMAZON_ALBUM::B0D6WXQCVL":
        {
            "id": "B0D6WXQCVL",
            "type": "album",
            "title": "Facing The End",
            "artistName": "APOCRYPHAL",
            "thumbnailUrl": "https://m.media-amazon.com/images/I/51QUsa2NDBL.jpg",
            "thumbnailWidth": 500,
            "thumbnailHeight": 500,
            "apiProvider": "amazon",
            "platforms":
            [
                "amazonMusic",
                "amazonStore"
            ]
        },
        "ANGHAMI_ALBUM::1053942249":
        {
            "id": "1053942249",
            "type": "album",
            "title": "Facing The End",
            "artistName": "Apocryphal",
            "thumbnailUrl": "https://artwork.anghcdn.co/cover/207874647/320",
            "thumbnailWidth": 1024,
            "thumbnailHeight": 1024,
            "apiProvider": "anghami",
            "platforms":
            [
                "anghami"
            ]
        },
        "BANDCAMP_ALBUM::2177048797":
        {
            "id": "2177048797",
            "type": "album",
            "title": "Facing the end",
            "artistName": "Apocryphal",
            "thumbnailUrl": "https://f4.bcbits.com/img/a3092362090_5.jpg",
            "thumbnailWidth": 700,
            "thumbnailHeight": 700,
            "apiProvider": "bandcamp",
            "platforms":
            [
                "bandcamp"
            ]
        },
        "BOOMPLAY_ALBUM::92209628":
        {
            "id": "92209628",
            "type": "album",
            "title": "Facing The End",
            "artistName": "Apocryphal",
            "thumbnailUrl": "https://source.boomplaymusic.com/group10/M00/08/07/496d3e0a17694fe7bd69247e86cec2c6H3000W3000_464_464.png",
            "thumbnailWidth": 464,
            "thumbnailHeight": 464,
            "apiProvider": "boomplay",
            "platforms":
            [
                "boomplay"
            ]
        },
        "ITUNES_ALBUM::1751932130":
        {
            "id": "1751932130",
            "type": "album",
            "title": "Facing the End",
            "artistName": "Apocryphal",
            "thumbnailUrl": "https://is1-ssl.mzstatic.com/image/thumb/Music211/v4/12/19/8d/12198d06-ec37-d1b8-3524-06ac9f2c1568/8721093338598.png/512x512bb.jpg",
            "thumbnailWidth": 512,
            "thumbnailHeight": 512,
            "apiProvider": "itunes",
            "platforms":
            [
                "appleMusic",
                "itunes"
            ]
        },
        "PANDORA_ALBUM::AL:35838569":
        {
            "id": "AL:35838569",
            "type": "album",
            "title": "Facing The End",
            "artistName": "Apocryphal",
            "thumbnailUrl": "https://content-images.p-cdn.com/images/3a/69/9d/a9/5cd94415896fce0058344767/_500W_500H.jpg",
            "thumbnailWidth": 500,
            "thumbnailHeight": 500,
            "apiProvider": "pandora",
            "platforms":
            [
                "pandora"
            ]
        }
    },
    "linksByPlatform":
    {
        "amazonMusic":
        {
            "country": "US",
            "url": "https://music.amazon.com/albums/B0D6WXQCVL",
            "entityUniqueId": "AMAZON_ALBUM::B0D6WXQCVL"
        },
        "amazonStore":
        {
            "country": "US",
            "url": "https://amazon.com/dp/B0D6WXQCVL",
            "entityUniqueId": "AMAZON_ALBUM::B0D6WXQCVL"
        },
        "anghami":
        {
            "country": "US",
            "url": "https://play.anghami.com/album/1053942249?refer=linktree",
            "entityUniqueId": "ANGHAMI_ALBUM::1053942249"
        },
        "bandcamp":
        {
            "country": "US",
            "url": "https://apocryphal-official.bandcamp.com/album/facing-the-end",
            "entityUniqueId": "BANDCAMP_ALBUM::2177048797"
        },
        "boomplay":
        {
            "country": "US",
            "url": "https://www.boomplay.com/albums/92209628",
            "entityUniqueId": "BOOMPLAY_ALBUM::92209628"
        },
        "pandora":
        {
            "country": "US",
            "url": "https://www.pandora.com/AL:35838569",
            "entityUniqueId": "PANDORA_ALBUM::AL:35838569"
        },
        "appleMusic":
        {
            "country": "US",
            "url": "https://geo.music.apple.com/us/album/_/1751932130?mt=1&app=music&ls=1&at=1000lHKX&ct=api_http&itscg=30200&itsct=odsl_m",
            "nativeAppUriMobile": "music://itunes.apple.com/us/album/_/1751932130?mt=1&app=music&ls=1&at=1000lHKX&ct=api_uri_m&itscg=30200&itsct=odsl_m",
            "nativeAppUriDesktop": "itmss://itunes.apple.com/us/album/_/1751932130?mt=1&app=music&ls=1&at=1000lHKX&ct=api_uri_d&itscg=30200&itsct=odsl_m",
            "entityUniqueId": "ITUNES_ALBUM::1751932130"
        },
        "itunes":
        {
            "country": "US",
            "url": "https://geo.music.apple.com/us/album/_/1751932130?mt=1&app=itunes&ls=1&at=1000lHKX&ct=api_http&itscg=30200&itsct=odsl_m",
            "nativeAppUriMobile": "itmss://itunes.apple.com/us/album/_/1751932130?mt=1&app=itunes&ls=1&at=1000lHKX&ct=api_uri_m&itscg=30200&itsct=odsl_m",
            "nativeAppUriDesktop": "itmss://itunes.apple.com/us/album/_/1751932130?mt=1&app=itunes&ls=1&at=1000lHKX&ct=api_uri_d&itscg=30200&itsct=odsl_m",
            "entityUniqueId": "ITUNES_ALBUM::1751932130"
        }
    }
}`
	resp := &GetLinksResponse{}
	err := json.Unmarshal([]byte(data), resp)
	if err != nil {
		log.Fatal(err)
	}
	if resp.EntitiesByUniqueId["BANDCAMP_ALBUM::2177048797"].Id != "2177048797" {
		log.Fatal("Title is not correct")
	}
}

func TestGetLinksResponse_Unmarshall_IDIsInt(t *testing.T) {
	data := `{
    "entityUniqueId": "ITUNES_ALBUM::1751932130",
    "userCountry": "US",
    "pageUrl": "https://album.link/us/i/1751932130",
    "entitiesByUniqueId":
    {
        "AMAZON_ALBUM::B0D6WXQCVL":
        {
            "id": "B0D6WXQCVL",
            "type": "album",
            "title": "Facing The End",
            "artistName": "APOCRYPHAL",
            "thumbnailUrl": "https://m.media-amazon.com/images/I/51QUsa2NDBL.jpg",
            "thumbnailWidth": 500,
            "thumbnailHeight": 500,
            "apiProvider": "amazon",
            "platforms":
            [
                "amazonMusic",
                "amazonStore"
            ]
        },
        "ANGHAMI_ALBUM::1053942249":
        {
            "id": "1053942249",
            "type": "album",
            "title": "Facing The End",
            "artistName": "Apocryphal",
            "thumbnailUrl": "https://artwork.anghcdn.co/cover/207874647/320",
            "thumbnailWidth": 1024,
            "thumbnailHeight": 1024,
            "apiProvider": "anghami",
            "platforms":
            [
                "anghami"
            ]
        },
        "BANDCAMP_ALBUM::2177048797":
        {
            "id": 2177048797,
            "type": "album",
            "title": "Facing the end",
            "artistName": "Apocryphal",
            "thumbnailUrl": "https://f4.bcbits.com/img/a3092362090_5.jpg",
            "thumbnailWidth": 700,
            "thumbnailHeight": 700,
            "apiProvider": "bandcamp",
            "platforms":
            [
                "bandcamp"
            ]
        },
        "BOOMPLAY_ALBUM::92209628":
        {
            "id": "92209628",
            "type": "album",
            "title": "Facing The End",
            "artistName": "Apocryphal",
            "thumbnailUrl": "https://source.boomplaymusic.com/group10/M00/08/07/496d3e0a17694fe7bd69247e86cec2c6H3000W3000_464_464.png",
            "thumbnailWidth": 464,
            "thumbnailHeight": 464,
            "apiProvider": "boomplay",
            "platforms":
            [
                "boomplay"
            ]
        },
        "ITUNES_ALBUM::1751932130":
        {
            "id": "1751932130",
            "type": "album",
            "title": "Facing the End",
            "artistName": "Apocryphal",
            "thumbnailUrl": "https://is1-ssl.mzstatic.com/image/thumb/Music211/v4/12/19/8d/12198d06-ec37-d1b8-3524-06ac9f2c1568/8721093338598.png/512x512bb.jpg",
            "thumbnailWidth": 512,
            "thumbnailHeight": 512,
            "apiProvider": "itunes",
            "platforms":
            [
                "appleMusic",
                "itunes"
            ]
        },
        "PANDORA_ALBUM::AL:35838569":
        {
            "id": "AL:35838569",
            "type": "album",
            "title": "Facing The End",
            "artistName": "Apocryphal",
            "thumbnailUrl": "https://content-images.p-cdn.com/images/3a/69/9d/a9/5cd94415896fce0058344767/_500W_500H.jpg",
            "thumbnailWidth": 500,
            "thumbnailHeight": 500,
            "apiProvider": "pandora",
            "platforms":
            [
                "pandora"
            ]
        }
    },
    "linksByPlatform":
    {
        "amazonMusic":
        {
            "country": "US",
            "url": "https://music.amazon.com/albums/B0D6WXQCVL",
            "entityUniqueId": "AMAZON_ALBUM::B0D6WXQCVL"
        },
        "amazonStore":
        {
            "country": "US",
            "url": "https://amazon.com/dp/B0D6WXQCVL",
            "entityUniqueId": "AMAZON_ALBUM::B0D6WXQCVL"
        },
        "anghami":
        {
            "country": "US",
            "url": "https://play.anghami.com/album/1053942249?refer=linktree",
            "entityUniqueId": "ANGHAMI_ALBUM::1053942249"
        },
        "bandcamp":
        {
            "country": "US",
            "url": "https://apocryphal-official.bandcamp.com/album/facing-the-end",
            "entityUniqueId": "BANDCAMP_ALBUM::2177048797"
        },
        "boomplay":
        {
            "country": "US",
            "url": "https://www.boomplay.com/albums/92209628",
            "entityUniqueId": "BOOMPLAY_ALBUM::92209628"
        },
        "pandora":
        {
            "country": "US",
            "url": "https://www.pandora.com/AL:35838569",
            "entityUniqueId": "PANDORA_ALBUM::AL:35838569"
        },
        "appleMusic":
        {
            "country": "US",
            "url": "https://geo.music.apple.com/us/album/_/1751932130?mt=1&app=music&ls=1&at=1000lHKX&ct=api_http&itscg=30200&itsct=odsl_m",
            "nativeAppUriMobile": "music://itunes.apple.com/us/album/_/1751932130?mt=1&app=music&ls=1&at=1000lHKX&ct=api_uri_m&itscg=30200&itsct=odsl_m",
            "nativeAppUriDesktop": "itmss://itunes.apple.com/us/album/_/1751932130?mt=1&app=music&ls=1&at=1000lHKX&ct=api_uri_d&itscg=30200&itsct=odsl_m",
            "entityUniqueId": "ITUNES_ALBUM::1751932130"
        },
        "itunes":
        {
            "country": "US",
            "url": "https://geo.music.apple.com/us/album/_/1751932130?mt=1&app=itunes&ls=1&at=1000lHKX&ct=api_http&itscg=30200&itsct=odsl_m",
            "nativeAppUriMobile": "itmss://itunes.apple.com/us/album/_/1751932130?mt=1&app=itunes&ls=1&at=1000lHKX&ct=api_uri_m&itscg=30200&itsct=odsl_m",
            "nativeAppUriDesktop": "itmss://itunes.apple.com/us/album/_/1751932130?mt=1&app=itunes&ls=1&at=1000lHKX&ct=api_uri_d&itscg=30200&itsct=odsl_m",
            "entityUniqueId": "ITUNES_ALBUM::1751932130"
        }
    }
}`
	resp := &GetLinksResponse{}
	err := json.Unmarshal([]byte(data), resp)
	if err != nil {
		log.Fatal(err)
	}
	if resp.EntitiesByUniqueId["BANDCAMP_ALBUM::2177048797"].Id != "2177048797" {
		log.Fatal("Title is not correct")
	}
}
