package odesli

type Platform string

const (
	PlatformAmazon       Platform = "amazon"
	PlatformAmazonMusic  Platform = "amazonMusic"
	PlatformAmazonStore  Platform = "amazonStore"
	PlatformAnghami      Platform = "anghami"
	PlatformAppleMusic   Platform = "appleMusic"
	PlatformAudiomack    Platform = "audiomack"
	PlatformAudius       Platform = "audius"
	PlatformBandcamp     Platform = "bandcamp"
	PlatformBoomplay     Platform = "boomplay"
	PlatformDeezer       Platform = "deezer"
	PlatformGoogle       Platform = "google"
	PlatformItunes       Platform = "itunes"
	PlatformNapster      Platform = "napster"
	PlatformPandora      Platform = "pandora"
	PlatformSoundcloud   Platform = "soundcloud"
	PlatformSpinrilla    Platform = "spinrilla"
	PlatformSpotify      Platform = "spotify"
	PlatformTidal        Platform = "tidal"
	PlatformYandex       Platform = "yandex"
	PlatformYoutube      Platform = "youtube"
	PlatformYoutubeMusic Platform = "youtubeMusic"
)

func (a Platform) String() string {
	return string(a)
}

func AvailablePlatforms() []Platform {
	return []Platform{
		PlatformAmazon,
		PlatformAmazonMusic,
		PlatformAmazonStore,
		PlatformAnghami,
		PlatformAppleMusic,
		PlatformAudiomack,
		PlatformAudius,
		PlatformBandcamp,
		PlatformBoomplay,
		PlatformDeezer,
		PlatformGoogle,
		PlatformItunes,
		PlatformNapster,
		PlatformPandora,
		PlatformSoundcloud,
		PlatformSpinrilla,
		PlatformSpotify,
		PlatformTidal,
		PlatformYandex,
		PlatformYoutube,
		PlatformYoutubeMusic,
	}
}

func PlatformFromString(s string) Platform {
	switch s {
	case "amazon":
		return PlatformAmazon
	case "amazonMusic":
		return PlatformAmazonMusic
	case "amazonStore":
		return PlatformAmazonStore
	case "anghami":
		return PlatformAnghami
	case "appleMusic":
		return PlatformAppleMusic
	case "audiomack":
		return PlatformAudiomack
	case "audius":
		return PlatformAudius
	case "bandcamp":
		return PlatformBandcamp
	case "boomplay":
		return PlatformBoomplay
	case "deezer":
		return PlatformDeezer
	case "google":
		return PlatformGoogle
	case "itunes":
		return PlatformItunes
	case "napster":
		return PlatformNapster
	case "pandora":
		return PlatformPandora
	case "soundcloud":
		return PlatformSoundcloud
	case "spinrilla":
		return PlatformSpinrilla
	case "spotify":
		return PlatformSpotify
	case "tidal":
		return PlatformTidal
	case "yandex":
		return PlatformYandex
	case "youtube":
		return PlatformYoutube
	case "youtubeMusic":
		return PlatformYoutubeMusic
	default:
		return Platform("")
	}
}

func (a Platform) ReadableName() string {
	switch a {
	case PlatformAmazon:
		return "Amazon"
	case PlatformAmazonMusic:
		return "Amazon Music"
	case PlatformAmazonStore:
		return "Amazon Store"
	case PlatformAnghami:
		return "Anghami"
	case PlatformAppleMusic:
		return "Apple Music"
	case PlatformAudiomack:
		return "Audiomack"
	case PlatformAudius:
		return "Audius"
	case PlatformBandcamp:
		return "Bandcamp"
	case PlatformBoomplay:
		return "Boomplay"
	case PlatformDeezer:
		return "Deezer"
	case PlatformGoogle:
		return "Google"
	case PlatformItunes:
		return "iTunes"
	case PlatformNapster:
		return "Napster"
	case PlatformPandora:
		return "Pandora"
	case PlatformSoundcloud:
		return "SoundCloud"
	case PlatformSpinrilla:
		return "Spinrilla"
	case PlatformSpotify:
		return "Spotify"
	case PlatformTidal:
		return "Tidal"
	case PlatformYandex:
		return "Yandex"
	case PlatformYoutube:
		return "YouTube"
	case PlatformYoutubeMusic:
		return "YouTube Music"
	default:
		return "Unknown"
	}
}
