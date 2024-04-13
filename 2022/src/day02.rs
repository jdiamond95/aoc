fn get_turn_value(my_move: char) -> u8 {
    match my_move {
        'X' => 1,
        'Y' => 2,
        'Z' => 3,
        _ => 0,
    }
}

fn what_did_i_pick(opponent_move: char, outcome: char) -> char {
    match (opponent_move, outcome) {
        ('A', 'X') => 'Z',
        ('A', 'Y') => 'X',
        ('A', 'Z') => 'Y',
        ('B', 'X') => 'X',
        ('B', 'Y') => 'Y',
        ('B', 'Z') => 'Z',
        ('C', 'X') => 'Y',
        ('C', 'Y') => 'Z',
        ('C', 'Z') => 'X',
        _ => opponent_move,
    }
}

fn get_score(opponent_move: char, my_move: char, part_one: bool) -> u8 {
    let outcome_score: u8 = match (opponent_move, my_move) {
        ('A', 'Y') => 6,
        ('A', 'Z') => 0,
        ('B', 'X') => 0,
        ('B', 'Z') => 6,
        ('C', 'X') => 6,
        ('C', 'Y') => 0,
        _ => 3,
    };
    if part_one {
        return get_turn_value(my_move) + outcome_score;
    } else {
        let outcome_score = match my_move {
            'X' => 0,
            'Y' => 3,
            _ => 6,
        };
        return outcome_score + get_turn_value(what_did_i_pick(opponent_move, my_move));
    }
}

pub fn run(input: &str) {
    println!("\n\nDay 02");
    let games = input.lines();

    let mut total_score1: i32 = 0;
    let mut total_score2: i32 = 0;

    for game in games {
        total_score1 += i32::from(get_score(
            game.chars().nth(0).unwrap(),
            game.chars().nth(2).unwrap(),
            true,
        ));

        total_score2 += i32::from(get_score(
            game.chars().nth(0).unwrap(),
            game.chars().nth(2).unwrap(),
            false,
        ));
    }
    println!("Part 1: {}", total_score1);
    println!("Part 2: {}", total_score2);
}
