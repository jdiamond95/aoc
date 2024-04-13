use std::collections::HashSet;

pub fn run(input: &str) {
    println!("\n\nDay 03");
    let split_compartments = input
        .lines()
        .map(|rucksack| rucksack.as_bytes().split_at(rucksack.len() / 2));

    let mut total_priority: u32 = 0;
    for bag in split_compartments {
        let cmp1: HashSet<_> = bag.0.iter().collect();
        let cmp2: HashSet<_> = bag.1.iter().collect();
        let intersection: Vec<_> = cmp1.intersection(&cmp2).collect();
        total_priority += u32::from(translate_priority(**intersection[0]));
    }
    println!("{total_priority}");
}

fn translate_priority(letter: u8) -> u8 {
    let mut priority = letter.clone();
    match priority {
        65..=90 => priority -= 38,
        97..=122 => priority -= 96,
        _ => {}
    }
    priority
}
