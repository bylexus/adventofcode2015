use super::Day;
use alex_lib::read_lines;

#[derive(Debug)]
pub struct Day07 {
    input: Vec<String>,
}

impl Day07 {
    pub fn new() -> Day07 {
        Day07 { input: Vec::new() }
    }

    fn parse_input(&mut self) {}
}

impl Day for Day07 {
    fn day_nr(&self) -> String {
        String::from("07")
    }
    fn title(&self) -> String {
        String::from("xxxxxxxx")
    }

    fn prepare(&mut self) {
        let input = read_lines("data/day07.txt");
        // let input = read_lines("data/day07-test.txt");
        self.input = input;
        self.parse_input();
    }

    fn solve1(&mut self) -> String {
        let mut solution: u64 = 0;

        String::from(format!("{0}", solution))
    }

    fn solve2(&mut self) -> String {
        let mut solution: u64 = 0;

        String::from(format!("{0}", solution))
    }
}
