package mq

import (
	"context"
	. "middleware/mq/mw"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestMQSuite(t *testing.T) {
	suite.Run(t, new(MQSuite))
}

type MQSuite struct {
	suite.Suite
}

func (m *MQSuite) SetupTest() {

}

func (m *MQSuite) TestMW() {
	mqHandler := NewMQHandler(m.handlerMsg)
	mqHandler.Register(TimeCostMW, FilterMW, LoggerMW)
	err := mqHandler.Exec(context.Background(), "hello chain")
	m.Nil(err)
}

func (m *MQSuite) handlerMsg(ctx context.Context, msg string) error {
	return nil
}
