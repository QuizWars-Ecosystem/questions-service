package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	AdminActionsTotalCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "admin_actions_total",
			Help: "Number of admin actions",
		},
		[]string{"method"},
	)

	AdminForbittenActionsTotalCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "admin_forbitten_actions_total",
			Help: "Number of forbitten admin actions",
		},
		[]string{"method", "reason"},
	)
)

var CategoriesTotalCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "categories_total",
		Help: "Number of categories",
	},
)

var (
	QuestionsRequestsGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "questions_requests",
			Help: "Number of question requests",
		},
	)

	QuestionsTotalGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "questions_total",
			Help: "Number of questions",
		},
	)

	QuestionsOptionsTotalGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "questions_options_total",
			Help: "Number of question's options",
		},
	)
)

func Initialize() {
	prometheus.MustRegister(AdminActionsTotalCounter)
	prometheus.MustRegister(AdminForbittenActionsTotalCounter)

	prometheus.MustRegister(CategoriesTotalCounter)

	prometheus.MustRegister(QuestionsRequestsGauge)
	prometheus.MustRegister(QuestionsTotalGauge)
	prometheus.MustRegister(QuestionsOptionsTotalGauge)
}
