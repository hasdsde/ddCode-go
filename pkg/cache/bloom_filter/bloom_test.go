package bloomfilter

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-faker/faker/v4"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MainSuite struct {
	suite.Suite
	filter *BloomFilter
	ctx    context.Context
}

func Test_NetQualityMainSuite(t *testing.T) {
	ctx := context.Background()
	// 1w个样本, 万分之一的误差
	filter, err := NewBloom(ctx, 1000000, 0.0001)
	assert.NoError(t, err)
	s := &MainSuite{
		ctx:    ctx,
		filter: filter,
	}
	suite.Run(t, s)
}

func (s *MainSuite) BeforeTest(suiteName, testName string) {
	s.filter.ClearAll()
}

func (s *MainSuite) Test_AddAndTest() {
	convey.Convey("Test_AddAndTest", s.T(), func() {
		convey.Reset(func() {
			s.BeforeTest("MainSuite", "AddAndTest")
		})
		// 构造测试用例
		exampleList := make([]string, 0, 1000)
		for i := 0; i < 1000; i++ {
			key := faker.ChineseName()
			exampleList = append(exampleList, key)
			s.filter.Add([]byte(key))
		}
		convey.Convey("存在测试", func() {
			for _, key := range exampleList {
				isExit := s.filter.Test([]byte(key))
				convey.So(isExit, convey.ShouldBeTrue)
			}
		})
		convey.Convey("aaa", func() {
			s.filter.DownloadToFile(context.Background(), "./aa")
		})
		convey.Convey("bbbb", func() {
			filter, err := NewBloom(context.Background(), 1000000, 0.0001, LoadFileWithOption("aa"))
			assert.NoError(s.T(), err)
			fmt.Println(filter)
		})
		// convey.Convey("不存在测试", func() {
		// 	for i := 0; i < 1000; i++ {
		// 		key := faker.DomainName()
		// 		isExit := s.filter.Test([]byte(key))
		// 		convey.So(isExit, convey.ShouldBeFalse)
		// 	}
		// })
	})
}

func TestBloomSaveOrLoad(t *testing.T) {
	ctx := context.Background()
	filter, err := NewBloom(ctx, 1000000, 0.0001)
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 1000; i++ {
		key := faker.ChineseName()
		t.Log(key)
		filter.Add([]byte(key))
	}
	filter.Add([]byte("aa22"))
	err = filter.DownloadToFile(ctx, "./aa")
	if err != nil {
		t.Fatal(err)
	}
	filter2, err := NewBloom(ctx, 1000000, 0.0001)
	if err != nil {
		t.Fatal(err)
	}
	err = filter2.LoadByFile(ctx, "./aa")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(filter2.Test([]byte("aa22")))
	t.Log(filter2.Test([]byte("aa12")))
}
