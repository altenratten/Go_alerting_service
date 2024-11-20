package main

import (
    "errors"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "sort"
    "strings"
    "time"
)

var (
    ErrUndefinedQuestion     = errors.New("неизвестный вопрос")
    ErrInvalidQuestionFormat = errors.New("неверный формат вопроса")
)

var (
    Database = map[string]string{
        "Алексей": "Калининград",
        "Алина":   "Красноярск",
        "Артём":   "Владивосток",
        "Дима":    "Челябинск",
        "Егор":    "Пермь",
        "Коля":    "Красноярск",
        "Миша":    "Москва",
        "Петя":    "Михайловка",
        "Сергей":  "Омск",
        "Соня":    "Москва",
    }

    UTCOffset = map[string]int{
        "Владивосток":     10,
        "Волгоград":       3,
        "Воронеж":         3,
        "Екатеринбург":    5,
        "Казань":          3,
        "Калининград":     2,
        "Краснодар":       3,
        "Красноярск":      7,
        "Москва":          3,
        "Нижний Новгород": 3,
        "Новосибирск":     7,
        "Омск":            6,
        "Пермь":           5,
        "Ростов-на-Дону":  3,
        "Самара":          4,
        "Санкт-Петербург": 3,
        "Уфа":             5,
        "Челябинск":       5,
    }
)

func formatFriendsCount(friendsCount int) string {
    if friendsCount == 1 {
        return "1 друг"
    }
    if friendsCount >= 2 && friendsCount <= 4 {
        return fmt.Sprintf("%d друга", friendsCount)
    }
    return fmt.Sprintf("%d друзей", friendsCount)
}

func whatTime(city string) string {
    offset := UTCOffset[city]
    currentTime := time.Now().UTC()
    cityTime := currentTime.Add(time.Duration(offset) * time.Hour)
    return cityTime.Format("15:04")
}

func whatWeather(city string) (string, error) {
    baseURL := "https://bba4h86i8icpvi5ot97n.containers.yandexcloud.net/go/weather"
    params := url.Values{}
    params.Add("0", "")
    params.Add("M", "")
    params.Add(city, "")

    urlInstance, err := url.Parse(baseURL)
    if err != nil {
        return "", fmt.Errorf("сетевая ошибка: %v", err)
    }
    urlInstance.RawQuery = params.Encode()

    response, err := http.Get(urlInstance.String())
    if err != nil {
        return "", fmt.Errorf("сетевая ошибка: %v", err)
    }

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return "", fmt.Errorf("ошибка на сервере погоды: %v", err)
    }
    content := string(body)
    return content, nil
}

func askFriend(friendName, question string) (string, error) {
    friendCity, exists := Database[friendName]
    if !exists {
        return "", fmt.Errorf("у тебя нет друга по имени %s", friendName)
    }
    if question == "ты где?" {
        return fmt.Sprintf("%s в городе: %s", friendName, friendCity), nil
    }
    if question == "который час?" {
        _, exists := UTCOffset[friendCity]
        if !exists {
            return "", fmt.Errorf("не могу определить время в городе: %s", friendCity)
        }
        cityTime := whatTime(friendCity)
        return fmt.Sprintf("Там сейчас %s", cityTime), nil
    }
	if question == "как погода?" {
		return whatWeather(friendCity)
	}
    // спроси друга о погоде

    return "", ErrUndefinedQuestion
}

func askAlice(question string) (string, error) {
    switch question {
    case "сколько у меня друзей?":
        friendsCount := len(Database)
        humanizedAnswer := formatFriendsCount(friendsCount)
        return fmt.Sprintf("У тебя %s", humanizedAnswer), nil
    case "кто все мои друзья?":
        friendsNames := make([]string, 0, len(Database))
        for friend := range Database {
            friendsNames = append(friendsNames, friend)
        }
        sort.Strings(friendsNames)
        friendsString := strings.Join(friendsNames, ", ")
        return fmt.Sprintf("Твои друзья: %s", friendsString), nil
    case "где все мои друзья?":
        friendsCities := make([]string, 0, len(Database))
        for _, city := range Database {
            friendsCities = append(friendsCities, city)
        }
        sort.Strings(friendsCities)
        citiesString := strings.Join(friendsCities, ", ")
        return fmt.Sprintf("Твои друзья в городах: %s", citiesString), nil
    default:
        return "", ErrUndefinedQuestion
    }
}

func askQuestion(question string) (string, error) {
    queryParts := strings.Split(question, ", ")
    if len(queryParts) != 2 {
        return "", ErrInvalidQuestionFormat
    }
    if queryParts[0] == "Алиса" {
        return askAlice(queryParts[1])
    }
    return askFriend(queryParts[0], queryParts[1])
}

func main() {
    questions := []string{
        "Алексей, который час?",
        "Алиса, где все мои друзья?",
        "Алиса, кто виноват?",
        "Алиса, кто все мои друзья?",
        "Алиса, сколько у меня друзей?",
        "Антон, как погода?",
        "Антон, который час?",
        "Антон, ты где?",
        "Артём, который час?",
        "Коля, как погода?",
        "Коля, ты где?",
        "Петя, который час?",
        "Соня, как погода?",
        "Соня, что делать?",
    }
    for _, question := range questions {
        answer, err := askQuestion(question)
        if err != nil {
            fmt.Println(question, "-", err)
            continue
        }
        fmt.Println(question, "-", answer)
    }
}