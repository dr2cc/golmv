package lmcz

import (
	"io"
	"net/http"
	"os"

	"github.com/dr2cc/golmv/internal/config"
	"github.com/dr2cc/golmv/internal/report"
)

// CzClient инкапсулирует методы настройки для работы с удаленным сервером.
// Хранение http.CzClient внутри структуры — это Go-идиома, позволяющая переиспользовать соединения.
type CzClient struct {
	cfg        config.LmczConfig
	httpClient *http.Client
	// [GUI-рефакторинг] out — куда клиент пишет сообщения о ходе работы.
	// Раньше методы печатали напрямую через fmt.Printf (то есть всегда в консоль).
	// Теперь печать идет через fmt.Fprintf(c.out, ...): для CLI это по-прежнему os.Stdout,
	// а GUI подставляет свой writer и показывает те же строки в окне лога.
	out io.Writer

	// [GUI-статусы] onStep получает изменения в состоянии шагов (шаги это find (поиск сервера) / status (получение статуса) / token (отправка токена)):
	// В CLI не задается (nil) — тогда "шаги" никуда не отправляются.
	//
	// Тип значения report.Func означает, что это функция, которая принимает один аргумент типа Step (снимок состояния шага) и ничего не возвращает
	// type Step struct {
	// 	ID     string
	// 	Title  string
	// 	State  State  // есть четыре состояния (state) - "pending", "running", "ok", "warn", "fail", "skipped"
	// 	Detail string `json:"detail"`
	// }
	onStep report.Func
}

// New создает и возвращает настроенный экземпляр Client
func NewCzClient(lmcz config.LmczConfig, httpClient *http.Client) *CzClient {
	return &CzClient{
		cfg:        lmcz,
		httpClient: httpClient,
		// [GUI-рефакторинг] По умолчанию — стандартный вывод, поэтому CLI работает как раньше,
		// а сигнатура конструктора не изменилась.
		out: os.Stdout,
	}
}
