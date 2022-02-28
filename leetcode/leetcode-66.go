<<<<<<< HEAD
func plusOne(digits []int) []int {
    for i:=len(digits)-1 ;i >=0 ; i -- {
        if digits[i] < 9{
            digits[i]++
            return digits
            }else{
            digits[i]=0
            
        }
    } 
    return append([]int{1},digits...)
=======
func plusOne(digits []int) []int {
    for i:=len(digits)-1 ;i >=0 ; i -- {
        if digits[i] < 9{
            digits[i]++
            return digits
            }else{
            digits[i]=0
            
        }
    } 
    return append([]int{1},digits...)
>>>>>>> 3101daa1d3c34d250c063e0fc3af2a0ec6cc4302
}