package pipeparser

import "testing"

func TestParse(t *testing.T) {
	testingData := []struct {
		input  string
		output []string
	}{
		{"cat test.txt", []string{"cat", "test.txt"}},
		{"ls -l -a", []string{"ls", "-l", "-a"}},
		{"ls -la", []string{"ls", "-la"}},
		{"awk '{print $1}'", []string{"awk", "{print $1}"}},
		{"ps -u test", []string{"ps", "-u", "test"}},
	}

	for _, data := range testingData {
		results := parse(data.input)

		for i := 0; i < len(results); i++ {
			if results[i] != data.output[i] {
				t.Errorf("Unmached result, input: %s, expected: %s, result: %s", data.input, data.output, results)
			}
		}
	}
}

func TestBuild(t *testing.T) {
	testingData := []struct {
		input       string
		expectedLen int
	}{
		{"cat test.txt", 1},
		{"ls -l -a", 1},
		{"ls -la", 1},
		{"awk '{print $1}'", 1},
		{"ps -u test", 1},
		{"cat test.txt | grep test", 2},
		{"cat test.txt | grep test | grep test2", 3},
		{"cat test.txt | grep test | grep test2 | awk '{print $1}'", 4},
	}

	for _, data := range testingData {
		results := buildCommands(data.input)

		if len(results) != data.expectedLen {
			t.Errorf("Unmached number of commands: input: %s, expected: %d, result: %d", data.input, data.expectedLen, len(results))
		}
	}
}

func TestRun(t *testing.T) {
	testingData := []struct {
		input     string
		hasErrors bool
	}{
		{"notarealcommand", true},
	}

	for _, data := range testingData {
		result, err := Run(data.input)

		if data.hasErrors && err == nil {
			t.Errorf("An error was expected: input: %s, result: %d", data.input, result)
		}

		if !data.hasErrors && err != nil {
			t.Errorf("An error occured: input: %s, result: %d, err: %s", data.input, result, err)
		}
	}
}
