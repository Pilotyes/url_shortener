# Курсовой проект url_shortener
## Проект url_shortener - это сервис коротких ссылок.
Позволяет пользователю из длинной ссылки получить короткую.
Предоставляет два интерфейса 1) HTTP 2) REST API
Чтобы получить короткую ссылку через HTTP интерфейс пользователь приходит на главную страницу сайта с ***запросом POST***, где в POST данных передает длинную ссылку - в ответ получает страницу, содержащую короткую ссылку.
Чтобы получить длинную ссылку пользователь приходит с запрсом на адрес короткой ссылки. Система отправляет ответ с установленным ***статус кодом 302*** и заголоком ***location*** с длинной ссылкой.
Также предусмотрен интерфейс ***REST API***, располагающийся по адресу **/api/v1**. При отправке на этот адрес запроса с методом POST и длинной ссылкой, возвращается JSON строка с короткой ссылкой и ошибкой, если она возникла.