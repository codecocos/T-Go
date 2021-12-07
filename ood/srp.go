package main

type FinanceReport struct {
}

func (r *FinanceReport) MakeReport() *Report {
	// ...
}

// //두 가지 책임을 가지고 있었서 샌드 방식이 바뀔 때 마다 변경 해주어야 하는 문제점이 있었음.
// 확장에는 닫혀있고 변경에는 열려있는 코드.
// func (r *FinanceReport) SendReport(method int) {
// 	switch (method) {
// 	case 1:
// 		//send email
// 	case 2:
// 		//make file
// 	case 3:
// 		// http
// 	case 4:
// 		// network
// 	}
// }

type ReportSender interface {
	SendReport(*Report)
}

type EmailReportSender struct {
}

func (s *EmailReportSender) SendReport(r *Report) {
	// .. email
}

type FileReportSender struct {
}

func (s *FileReportSender) SendReport(r *Report) {
	// .. make file
}

type HttpReportSender struct {
}

func (s *HttpReportSender) SendReport(r *Report) {
	// .. make file
}
