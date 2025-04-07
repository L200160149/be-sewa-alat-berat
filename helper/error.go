package helper

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
// func CheckError(err error) error {
//     if err != nil {
//         return err
//     }
//     return nil
// }