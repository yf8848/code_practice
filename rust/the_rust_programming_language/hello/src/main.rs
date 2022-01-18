fn main() {
    let num = 3;
    if num < 5 {
        print!("condition is true\n")
    } else {
        print!("condition is false\n")
    }

    let condition = true;
    let num = if condition { 5 } else { 50 };
    print!("the value of num is {}\n", num);

    let mut counter = 0;
    let num = loop {
        counter += 1;
        if counter == 10 {
            break counter * 2;
        }
    };
    print!("the value of loop num is {}\n", num)
}
