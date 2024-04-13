pub fn run(input: &str) {
    println!("\n\nDay 01");
    let mut calories_list: Vec<i32> = input
        .split("\n\n")
        .map(|elf| {
            elf.lines()
                .map(|food| food.parse::<i32>().unwrap())
                .sum::<i32>()
        })
        .collect();

    calories_list.sort();
    println!("Part 1: {}", calories_list.last().unwrap());

    println!(
        "Part 2: {}",
        calories_list.iter().rev().take(3).sum::<i32>()
    );
}
