package bean

import (
	"log"

	"github.com/natepboat/go-rest-api/api"
	"github.com/natepboat/go-rest-api/interceptor"
)

type BeanContext struct {
	beanMap map[string]interface{}
}

func (b *BeanContext) RequiredBean(name string) interface{} {
	bean, exist := b.beanMap[name]

	if !exist {
		log.Fatalf("Required %s but does not exist\n", name)
	}

	return bean
}

func InitBeanContext() *BeanContext {
	beanContext := &BeanContext{
		beanMap: make(map[string]interface{}, 5),
	}

	beanContext.beanMap[UserService] = CreateUserService()
	beanContext.beanMap[AuthInterceptor] = interceptor.NewAuthInterceptor()
	beanContext.beanMap[MetricInterceptor] = interceptor.NewMetricInterceptor()
	beanContext.beanMap[StatController] = api.NewStatController()
	beanContext.beanMap[UserController] = CreateUserController(beanContext)

	return beanContext
}
