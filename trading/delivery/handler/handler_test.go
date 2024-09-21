package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"tokenize/domain"
	"tokenize/domain/mocks"
	tradingHandler "tokenize/trading/delivery/handler"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TradingHandlerSuite struct {
	suite.Suite
	tradingUsecase *mocks.TradingUsecase
	engine         *gin.Engine
	router         *gin.RouterGroup
	w              *httptest.ResponseRecorder
}

func TestRun(t *testing.T) {
	suite.Run(t, &TradingHandlerSuite{})
}

func (s *TradingHandlerSuite) SetupTest() {
	s.tradingUsecase = new(mocks.TradingUsecase)
	s.engine = gin.Default()
	s.router = s.engine.Group("/task")
	s.w = httptest.NewRecorder()
}

func (s *TradingHandlerSuite) TestOrderBookGen_Success() {
	orderBook := &domain.OrderBook{
		Bids: []domain.Order{
			{Qty: 63.7974, Price: 0.04062, Sum: 2.591450388}, {Qty: 64.2797, Price: 0.04061, Sum: 2.6103986170000004}, {Qty: 45.6883, Price: 0.0406, Sum: 1.8549449799999997}, {Qty: 61.0201, Price: 0.04059, Sum: 2.476805859}, {Qty: 21.9508, Price: 0.04058, Sum: 0.890763464}, {Qty: 32.6104, Price: 0.04057, Sum: 1.323003928}, {Qty: 18.1411, Price: 0.04056, Sum: 0.7358030160000001}, {Qty: 17.1511, Price: 0.04055, Sum: 0.695477105}, {Qty: 33.2434, Price: 0.04054, Sum: 1.347687436}, {Qty: 25.4596, Price: 0.04053, Sum: 1.0318775879999997}, {Qty: 37.7625, Price: 0.04052, Sum: 1.5301365000000002}, {Qty: 4.1657, Price: 0.04051, Sum: 0.168752507}, {Qty: 4.1667, Price: 0.0405, Sum: 0.16875135}, {Qty: 82.4535, Price: 0.04049, Sum: 3.338542215}, {Qty: 6.0717, Price: 0.04048, Sum: 0.245782416}, {Qty: 6.9199, Price: 0.04047, Sum: 0.28004835299999997}, {Qty: 8.4245, Price: 0.04046, Sum: 0.34085527000000004}, {Qty: 102.2267, Price: 0.04045, Sum: 4.135070015}, {Qty: 4.1991, Price: 0.04044, Sum: 0.16981160399999998}, {Qty: 18.3719, Price: 0.04043, Sum: 0.742775917}, {Qty: 22.3333, Price: 0.04042, Sum: 0.902711986}, {Qty: 4.47, Price: 0.04041, Sum: 0.1806327}, {Qty: 87.389, Price: 0.0404, Sum: 3.5305155999999998}, {Qty: 26.4864, Price: 0.04039, Sum: 1.069785696}, {Qty: 5.289, Price: 0.04038, Sum: 0.21356982}, {Qty: 2.7548, Price: 0.04037, Sum: 0.11121127600000001}, {Qty: 2.9744, Price: 0.04036, Sum: 0.120046784}, {Qty: 3.3044, Price: 0.04035, Sum: 0.13333253999999997}, {Qty: 3.0038, Price: 0.04034, Sum: 0.121173292}, {Qty: 3.2187, Price: 0.04033, Sum: 0.129810171}, {Qty: 3.0592, Price: 0.04032, Sum: 0.12334694400000001}, {Qty: 3.0361, Price: 0.04031, Sum: 0.12238519099999999}, {Qty: 18.1893, Price: 0.0403, Sum: 0.73302879}, {Qty: 2.7075, Price: 0.04029, Sum: 0.10908517499999999}, {Qty: 3.1975, Price: 0.04028, Sum: 0.1287953}, {Qty: 3.5068, Price: 0.04027, Sum: 0.14121883600000001}, {Qty: 14.708, Price: 0.04026, Sum: 0.59214408}, {Qty: 6.9223, Price: 0.04025, Sum: 0.278622575}, {Qty: 4.4639, Price: 0.04024, Sum: 0.17962733599999997}, {Qty: 4.2978, Price: 0.04023, Sum: 0.172900494}, {Qty: 3.457, Price: 0.04022, Sum: 0.13904054}, {Qty: 2.973, Price: 0.04021, Sum: 0.11954433}, {Qty: 133.5885, Price: 0.0402, Sum: 5.370257700000001}, {Qty: 6.7405, Price: 0.04019, Sum: 0.27090069499999997}, {Qty: 6.2233, Price: 0.04018, Sum: 0.250052194}, {Qty: 3.1488, Price: 0.04017, Sum: 0.126487296}, {Qty: 3.2465, Price: 0.04016, Sum: 0.13037944}, {Qty: 3.483, Price: 0.04015, Sum: 0.13984245}, {Qty: 2.8202, Price: 0.04014, Sum: 0.113202828}, {Qty: 4.5308, Price: 0.04013, Sum: 0.181821004}, {Qty: 4.3382, Price: 0.04012, Sum: 0.174048584}, {Qty: 3.2218, Price: 0.04011, Sum: 0.129226398}, {Qty: 5.8074, Price: 0.0401, Sum: 0.23287674}, {Qty: 3.2807, Price: 0.04009, Sum: 0.131523263}, {Qty: 4.2329, Price: 0.04008, Sum: 0.16965463199999997}, {Qty: 3.1459, Price: 0.04007, Sum: 0.126056213}, {Qty: 3.147, Price: 0.04006, Sum: 0.12606882}, {Qty: 51.0002, Price: 0.04005, Sum: 2.04255801}, {Qty: 42.4076, Price: 0.04004, Sum: 1.698000304}, {Qty: 2.6584, Price: 0.04003, Sum: 0.106415752}, {Qty: 2.8574, Price: 0.04002, Sum: 0.114353148}, {Qty: 3.8163, Price: 0.04001, Sum: 0.152690163}, {Qty: 30.4305, Price: 0.04, Sum: 1.21722}, {Qty: 2.5698, Price: 0.03999, Sum: 0.10276630199999999}, {Qty: 2.9268, Price: 0.03998, Sum: 0.11701346400000001}, {Qty: 3.266, Price: 0.03997, Sum: 0.13054202}, {Qty: 13.9415, Price: 0.03996, Sum: 0.55710234}, {Qty: 2.6458, Price: 0.03995, Sum: 0.10569971}, {Qty: 6.0204, Price: 0.03994, Sum: 0.24045477600000004}, {Qty: 3.3829, Price: 0.03993, Sum: 0.13507919699999998}, {Qty: 2.6774, Price: 0.03992, Sum: 0.106881808}, {Qty: 3.286, Price: 0.03991, Sum: 0.13114426}, {Qty: 7.2189, Price: 0.0399, Sum: 0.28803410999999995}, {Qty: 26.8432, Price: 0.03989, Sum: 1.0707752480000001}, {Qty: 16.0676, Price: 0.03988, Sum: 0.640775888}, {Qty: 4.4375, Price: 0.03987, Sum: 0.17692312500000001}, {Qty: 3.2534, Price: 0.03986, Sum: 0.129680524}, {Qty: 85.871, Price: 0.03985, Sum: 3.4219593499999994}, {Qty: 4.3378, Price: 0.03984, Sum: 0.172817952}, {Qty: 2.7771, Price: 0.03983, Sum: 0.11061189299999999}, {Qty: 2.7367, Price: 0.03982, Sum: 0.108975394}, {Qty: 4.2149, Price: 0.03981, Sum: 0.167795169}, {Qty: 14.27, Price: 0.0398, Sum: 0.5679460000000001}, {Qty: 15.5716, Price: 0.03979, Sum: 0.619593964}, {Qty: 3.8574, Price: 0.03978, Sum: 0.15344737200000003}, {Qty: 3.0351, Price: 0.03977, Sum: 0.12070592699999999}, {Qty: 2.9062, Price: 0.03976, Sum: 0.115550512}, {Qty: 2.9721, Price: 0.03975, Sum: 0.11814097500000001}, {Qty: 3.4837, Price: 0.03974, Sum: 0.138442238}, {Qty: 6.6039, Price: 0.03973, Sum: 0.262372947}, {Qty: 2.9008, Price: 0.03972, Sum: 0.11521977599999998}, {Qty: 4.5225, Price: 0.03971, Sum: 0.179588475}, {Qty: 5.0675, Price: 0.0397, Sum: 0.20117975}, {Qty: 1.1763, Price: 0.03969, Sum: 0.046687347}, {Qty: 25.547, Price: 0.03968, Sum: 1.01370496}, {Qty: 0.2595, Price: 0.03967, Sum: 0.010294365}, {Qty: 0.3534, Price: 0.03966, Sum: 0.014015844}, {Qty: 0.2036, Price: 0.03965, Sum: 0.00807274}, {Qty: 0.9276, Price: 0.03964, Sum: 0.036770064}, {Qty: 1.3312, Price: 0.03963, Sum: 0.05275545599999999},
		},
		Asks: []domain.Order{
			{Qty: 21.2278, Price: 0.04063, Sum: 0.862485514}, {Qty: 36.2771, Price: 0.04064, Sum: 1.474301344}, {Qty: 43.7293, Price: 0.04065, Sum: 1.7775960450000001}, {Qty: 30.4375, Price: 0.04066, Sum: 1.23758875}, {Qty: 50.9135, Price: 0.04067, Sum: 2.0706520449999997}, {Qty: 33.378, Price: 0.04068, Sum: 1.35781704}, {Qty: 42.3804, Price: 0.04069, Sum: 1.724458476}, {Qty: 46.0367, Price: 0.0407, Sum: 1.87369369}, {Qty: 9.2868, Price: 0.04071, Sum: 0.378065628}, {Qty: 53.855, Price: 0.04072, Sum: 2.1929756}, {Qty: 26.1181, Price: 0.04073, Sum: 1.063790213}, {Qty: 10.1731, Price: 0.04074, Sum: 0.414452094}, {Qty: 16.9174, Price: 0.04075, Sum: 0.68938405}, {Qty: 81.3418, Price: 0.04076, Sum: 3.3154917680000002}, {Qty: 22.6755, Price: 0.04077, Sum: 0.924480135}, {Qty: 2.736, Price: 0.04078, Sum: 0.11157408}, {Qty: 3.3778, Price: 0.04079, Sum: 0.137780462}, {Qty: 81.349, Price: 0.0408, Sum: 3.3190392}, {Qty: 3.2141, Price: 0.04081, Sum: 0.131167421}, {Qty: 3.7633, Price: 0.04082, Sum: 0.153617906}, {Qty: 3.1478, Price: 0.04083, Sum: 0.128524674}, {Qty: 3.4732, Price: 0.04084, Sum: 0.141845488}, {Qty: 26.8266, Price: 0.04085, Sum: 1.0958666099999999}, {Qty: 5.9422, Price: 0.04086, Sum: 0.242798292}, {Qty: 3.1574, Price: 0.04087, Sum: 0.129042938}, {Qty: 7.3281, Price: 0.04088, Sum: 0.299572728}, {Qty: 2.8406, Price: 0.04089, Sum: 0.116152134}, {Qty: 5.2068, Price: 0.0409, Sum: 0.21295812}, {Qty: 3.1607, Price: 0.04091, Sum: 0.129304237}, {Qty: 19.8723, Price: 0.04092, Sum: 0.8131745159999999}, {Qty: 126.0333, Price: 0.04093, Sum: 5.158542969}, {Qty: 4.4391, Price: 0.04094, Sum: 0.18173675399999997}, {Qty: 9.0282, Price: 0.04095, Sum: 0.36970479}, {Qty: 4.7703, Price: 0.04096, Sum: 0.195391488}, {Qty: 3.1629, Price: 0.04097, Sum: 0.129584013}, {Qty: 22.8976, Price: 0.04098, Sum: 0.938343648}, {Qty: 3.9759, Price: 0.04099, Sum: 0.16297214100000001}, {Qty: 51.2049, Price: 0.041, Sum: 2.0994009}, {Qty: 4.6312, Price: 0.04101, Sum: 0.189925512}, {Qty: 3.299, Price: 0.04102, Sum: 0.13532498}, {Qty: 3.4419, Price: 0.04103, Sum: 0.14122115699999999}, {Qty: 6.8863, Price: 0.04104, Sum: 0.282613752}, {Qty: 3.1379, Price: 0.04105, Sum: 0.128810795}, {Qty: 3.8791, Price: 0.04106, Sum: 0.159275846}, {Qty: 3.0439, Price: 0.04107, Sum: 0.125012973}, {Qty: 4.5552, Price: 0.04108, Sum: 0.187127616}, {Qty: 43.3617, Price: 0.04109, Sum: 1.781732253}, {Qty: 9.0746, Price: 0.0411, Sum: 0.37296606}, {Qty: 48.3285, Price: 0.04111, Sum: 1.986784635}, {Qty: 4.4585, Price: 0.04112, Sum: 0.18333351999999997}, {Qty: 3.0579, Price: 0.04113, Sum: 0.125771427}, {Qty: 3.803, Price: 0.04114, Sum: 0.15645542}, {Qty: 5.5384, Price: 0.04115, Sum: 0.22790516}, {Qty: 8.6159, Price: 0.04116, Sum: 0.354630444}, {Qty: 80.954, Price: 0.04117, Sum: 3.3328761799999995}, {Qty: 4.3335, Price: 0.04118, Sum: 0.17845353}, {Qty: 3.1316, Price: 0.04119, Sum: 0.128990604}, {Qty: 7.494, Price: 0.0412, Sum: 0.3087528}, {Qty: 3.9276, Price: 0.04121, Sum: 0.16185639599999999}, {Qty: 3.5, Price: 0.04122, Sum: 0.14427}, {Qty: 3.0318, Price: 0.04123, Sum: 0.12500111400000002}, {Qty: 4.4963, Price: 0.04124, Sum: 0.18542741199999999}, {Qty: 4.2191, Price: 0.04125, Sum: 0.174037875}, {Qty: 2.9453, Price: 0.04126, Sum: 0.12152307799999999}, {Qty: 27.6098, Price: 0.04127, Sum: 1.139456446}, {Qty: 4.2973, Price: 0.04128, Sum: 0.17739254399999999}, {Qty: 3.1893, Price: 0.04129, Sum: 0.131686197}, {Qty: 29.4475, Price: 0.0413, Sum: 1.21618175}, {Qty: 3.5755, Price: 0.04131, Sum: 0.147703905}, {Qty: 5.2856, Price: 0.04132, Sum: 0.218400992}, {Qty: 3.5315, Price: 0.04133, Sum: 0.145956895}, {Qty: 6.5744, Price: 0.04134, Sum: 0.271785696}, {Qty: 3.6975, Price: 0.04135, Sum: 0.15289162499999998}, {Qty: 4.7189, Price: 0.04136, Sum: 0.195173704}, {Qty: 4.3166, Price: 0.04137, Sum: 0.17857774199999998}, {Qty: 3.3741, Price: 0.04138, Sum: 0.139620258}, {Qty: 3.0503, Price: 0.04139, Sum: 0.12625191700000002}, {Qty: 9.2965, Price: 0.0414, Sum: 0.38487509999999997}, {Qty: 2.7105, Price: 0.04141, Sum: 0.11224180500000001}, {Qty: 185.6512, Price: 0.04142, Sum: 7.6896727039999995}, {Qty: 2.8164, Price: 0.04143, Sum: 0.116683452}, {Qty: 4.1506, Price: 0.04144, Sum: 0.17200086399999998}, {Qty: 16.7036, Price: 0.04145, Sum: 0.6923642200000001}, {Qty: 3.0257, Price: 0.04146, Sum: 0.125445522}, {Qty: 3.0962, Price: 0.04147, Sum: 0.128399414}, {Qty: 4.2811, Price: 0.04148, Sum: 0.17758002800000003}, {Qty: 16.4071, Price: 0.04149, Sum: 0.680730579}, {Qty: 35.1525, Price: 0.0415, Sum: 1.4588287500000001}, {Qty: 0.528, Price: 0.04151, Sum: 0.02191728}, {Qty: 1.3534, Price: 0.04152, Sum: 0.056193168}, {Qty: 0.4024, Price: 0.04153, Sum: 0.016711671999999997}, {Qty: 0.9728, Price: 0.04154, Sum: 0.040410112}, {Qty: 0.8396, Price: 0.04155, Sum: 0.03488538}, {Qty: 1.9099, Price: 0.04156, Sum: 0.079375444}, {Qty: 4.1344, Price: 0.04157, Sum: 0.17186700800000002}, {Qty: 0.5673, Price: 0.04158, Sum: 0.023588334}, {Qty: 0.2112, Price: 0.04159, Sum: 0.008783808}, {Qty: 21.903, Price: 0.0416, Sum: 0.9111647999999999}, {Qty: 5.6246, Price: 0.04161, Sum: 0.234039606}, {Qty: 0.3098, Price: 0.04162, Sum: 0.012893876},
		},
	}

	orderBookMarshal, _ := json.Marshal(orderBook)
	s.tradingUsecase.On("GetOrderBook", mock.Anything, mock.MatchedBy(func(symbol string) bool { return true })).
		Return(orderBook, nil).Once()

	handler := tradingHandler.TradingHandler{
		TradingUsecase: s.tradingUsecase,
	}
	s.router.GET("", handler.OrderBookGen)
	req, _ := http.NewRequest("GET", "/task", nil)
	s.engine.ServeHTTP(s.w, req)

	assert.Equal(s.Suite.T(), http.StatusOK, s.w.Code)
	assert.Equal(s.Suite.T(), string(orderBookMarshal), s.w.Body.String())
}
