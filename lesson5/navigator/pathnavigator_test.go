package navigator

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"mtsbank_golang/lesson5/mock"
	"testing"
)

type PathNavigatorTestSuite struct {
	suite.Suite

	ctrl     *gomock.Controller
	navi     *PathNavigator
	mockPath *mock.MockPath
}

func TestPathNavigatorSuite(t *testing.T) {
	suite.Run(t, new(PathNavigatorTestSuite))
}

func (s *PathNavigatorTestSuite) SetupSuite() {
	s.ctrl = gomock.NewController(s.T())

	s.mockPath = mock.NewMockPath(s.ctrl)

	s.mockPath.EXPECT().Distance().Return(10.0, nil)

	s.mockPath.EXPECT().CountPoints().Return(4).AnyTimes()

	s.mockPath.EXPECT().DistanceBetween(-1, 0).Return(0.0, errors.New("my error")).AnyTimes()
	s.mockPath.EXPECT().DistanceBetween(0, 1).Return(2.0, nil).AnyTimes()
	s.mockPath.EXPECT().DistanceBetween(1, 2).Return(3.0, nil).AnyTimes()
	s.mockPath.EXPECT().DistanceBetween(2, 3).Return(5.0, nil).AnyTimes()
	s.mockPath.EXPECT().DistanceBetween(3, 4).Return(0.0, errors.New("my error")).AnyTimes()

	s.mockPath.EXPECT().PointAt(0).Return("p1", nil).AnyTimes()
	s.mockPath.EXPECT().PointAt(1).Return("p2", nil).AnyTimes()
	s.mockPath.EXPECT().PointAt(2).Return("p3", nil).AnyTimes()
	s.mockPath.EXPECT().PointAt(3).Return("p4", nil).AnyTimes()
}

func (s *PathNavigatorTestSuite) TearDown() {
	s.ctrl.Finish()
}

func (s *PathNavigatorTestSuite) TestPathNavigator_DistanceLeft() {
	s.navi, _ = NewPathNavigator(s.mockPath)
	distancesLeft := []float64{10, 8, 5, 0}
	var i = 0

	for s.navi.MoveNext() {
		_, _ = s.navi.CurrentLocation()
		assert.Equal(s.T(), s.navi.DistanceLeft(), distancesLeft[i])
		i++
	}
}

func (s *PathNavigatorTestSuite) TestPathNavigator_DistancePassed() {
	s.navi, _ = NewPathNavigator(s.mockPath)
	distancesPassed := []float64{0, 2, 5, 10}
	var i = 0

	for s.navi.MoveNext() {
		_, _ = s.navi.CurrentLocation()
		assert.Equal(s.T(), s.navi.DistancePassed(), distancesPassed[i])
		i++
	}
}

func (s *PathNavigatorTestSuite) TestPathNavigator_CurrentLocation() {
	s.navi, _ = NewPathNavigator(s.mockPath)
	locations := []string{"p1", "p2", "p3", "p4"}
	var i = 0

	for s.navi.MoveNext() {
		l, e := s.navi.CurrentLocation()

		if e != nil {
			s.T().Error(e)
		}
		s.Equal(l, locations[i])
		i++
	}
}

func (s *PathNavigatorTestSuite) TestPathNavigator_MoveNext() {
	s.navi, _ = NewPathNavigator(s.mockPath)
	er := []bool{true, true, true, true, false}

	for _, r := range er {
		s.Equal(r, s.navi.MoveNext())
	}
}
