package command

//func runActions(execAction IRunAction) error {
//
//	execAction.Construct()
//	if err := object.Set(execAction, os.Args); err != nil {
//		return fmt.Errorf("cmd set Args error: %w", err)
//	}
//	if err := execAction.Init(); err != nil {
//		return fmt.Errorf("cmd Init error: %w", err)
//	}
//	wait := make(chan bool, 1)
//	go func() {
//		execAction.Run()
//		execAction.Destruct()
//		wait <- true
//	}()
//	select {
//	case <-ctx.Done():
//	case <-wait:
//	}
//	return nil
//}
//
//func runApp() error {
//
//	//从app中找函数执行
//	if execApp != nil {
//		if method := object.FindMethod(execApp, nameCmd); method != "" {
//			wait := make(chan bool, 1)
//			go func() {
//				run := reflect.ValueOf(execApp).MethodByName(method)
//				run.Call([]reflect.Value{})
//				wait <- true
//			}()
//			select {
//			case <-ctx.Done():
//			case <-wait:
//			}
//			return nil
//		}
//	}
//	return errNoCmd
//}
