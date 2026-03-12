package errore

import (
	"errors"
	"fmt"
	"os"
)

// label - уникальное наименование
type label string

// command - команда, которую можно выполнять в игре
type command label

// список доступных команд
var (
	eat  = command("eat")
	take = command("take")
	talk = command("talk to")
)

// thing - объект, который существует в игре
type thing struct {
	name    label
	actions map[command]string
}

// supports() возвращает true, если объект
// поддерживает команду action
func (t thing) supports(action command) bool {
	_, ok := t.actions[action]
	return ok
}

// String() возвращает описание объекта
func (t thing) String() string {
	return string(t.name)
}

// полный список объектов в игре
var (
	apple = thing{"apple", map[command]string{
		eat:  "mmm, delicious!",
		take: "you have an apple now",
	}}
	bob = thing{"bob", map[command]string{
		talk: "Bob says hello",
	}}
	coin = thing{"coin", map[command]string{
		take: "you have a coin now",
	}}
	mirror = thing{"mirror", map[command]string{
		take: "you have a mirror now",
		talk: "mirror does not answer",
	}}
	mushroom = thing{"mushroom", map[command]string{
		eat:  "tastes funny",
		take: "you have a mushroom now",
	}}
)

// step описывает шаг игры: сочетание команды и объекта
type step struct {
	cmd command
	obj thing
}

// isValid() возвращает true, если объект
// совместим с командой
func (s step) isValid() bool {
	return s.obj.supports(s.cmd)
}

// String() возвращает описание шага
func (s step) String() string {
	return fmt.Sprintf("%s %s", s.cmd, s.obj)
}

// начало решения

// invalidStepError - ошибка, которая возникает,
// когда команда шага не совместима с объектом
type invalidStepError struct {
	cmd command
	obj thing
}

func (iSE invalidStepError) Error() string {
	return fmt.Sprintf("things like '%s %s' are impossible", iSE.cmd, iSE.obj.name)
}

// notEnoughObjectsError - ошибка, которая возникает,
// когда в игре закончились объекты определенного типа
type notEnoughObjectsError struct {
	obj thing
}

func (nEOE notEnoughObjectsError) Error() string {
	return fmt.Sprintf("be careful with scarce %ss", nEOE.obj.name)
}

// commandLimitExceededError - ошибка, которая возникает,
// когда игрок превысил лимит на выполнение команды
type commandLimitExceededError struct {
	cmd command
}

func (cLEE commandLimitExceededError) Error() string {
	return fmt.Sprintf("%s less", string(cLEE.cmd))
}

// objectLimitExceededError - ошибка, которая возникает,
// когда игрок превысил лимит на количество объектов
// определенного типа в инвентаре
type objectLimitExceededError struct {
	limit int
	obj   thing
}

func (oLEE objectLimitExceededError) Error() string {
	return fmt.Sprintf("don't be greedy, %d %s is enough", oLEE.limit, oLEE.obj.name)
}

// gameOverError - ошибка, которая произошла в игре
type gameOverError struct {
	// количество шагов, успешно выполненных
	// до того, как произошла ошибка
	nSteps int
	err    error
}

func (ge gameOverError) Error() string {
	return fmt.Sprintf("%v", ge.err)
}

func (ge gameOverError) Unwrap() error {
	return ge.err
}

// player - игрок
type player struct {
	// количество съеденного
	nEaten int
	// количество диалогов
	nDialogs int
	// инвентарь
	inventory []thing
}

// has() возвращает true, если у игрока
// в инвентаре есть предмет obj
func (p *player) has(obj thing) bool {
	for _, got := range p.inventory {
		if got.name == obj.name {
			return true
		}
	}
	return false
}

// do() выполняет команду cmd над объектом obj
// от имени игрока
func (p *player) do(cmd command, obj thing) error {
	// действуем в соответствии с командой
	switch cmd {
	case eat:
		if p.nEaten > 1 {
			return commandLimitExceededError{cmd} // ????
			// return errors.New("you don't want to eat anymore")
		}
		p.nEaten++
	case take:
		if p.has(obj) {
			return objectLimitExceededError{1, obj}
			// return fmt.Errorf("you already have a %s", obj)
		}
		p.inventory = append(p.inventory, obj)
	case talk:
		if p.nDialogs > 0 {
			return commandLimitExceededError{cmd}
			// return errors.New("you don't want to talk anymore")
		}
		p.nDialogs++
	}
	return nil
}

// newPlayer создает нового игрока
func newPlayer() *player {
	return &player{inventory: []thing{}}
}

// game описывает игру
type game struct {
	// игрок
	player *player
	// объекты игрового мира
	things map[label]int
	// количество успешно выполненных шагов
	nSteps int
}

// has() проверяет, остались ли в игровом мире указанные предметы
func (g *game) has(obj thing) bool {
	count := g.things[obj.name]
	return count > 0
}

// execute() выполняет шаг step
func (g *game) execute(st step) error {
	// проверяем совместимость команды и объекта
	if !st.isValid() {
		return gameOverError{g.nSteps, invalidStepError{st.cmd, st.obj}}
		// return fmt.Errorf("cannot %s", st)
		// return invalidStepError{st.cmd, st.obj}
	}

	// когда игрок берет или съедает предмет,
	// тот пропадает из игрового мира
	if st.cmd == take || st.cmd == eat {
		if !g.has(st.obj) {
			return gameOverError{g.nSteps, notEnoughObjectsError{st.obj}}
			// return fmt.Errorf("there are no %ss left", st.obj)
			// return notEnoughObjectsError{st.obj}

		}
		g.things[st.obj.name]--
	}

	// выполняем команду от имени игрока
	if err := g.player.do(st.cmd, st.obj); err != nil {
		return gameOverError{g.nSteps, err}
	}

	g.nSteps++
	return nil
}

// newGame() создает новую игру
func newGame() *game {
	p := newPlayer()
	things := map[label]int{
		apple.name:    2,
		coin.name:     3,
		mirror.name:   1,
		mushroom.name: 1,
	}
	return &game{p, things, 0}
}

// giveAdvice() возвращает совет, который
// поможет игроку избежать ошибки err в будущем
func giveAdvice(err error) string {
	var gameOverErrors gameOverError
	if errors.As(err, &gameOverErrors) {
		// if gameOverErrors.err == invalidStepError{} {
		// 	err = fmt.Errorf("wrap: %w", err)
		// }
		// fmt.Println("Language error:", gameOverErrors.err)
		var invalidStepErrors invalidStepError
		if errors.As(gameOverErrors.err, &invalidStepErrors) {
			return fmt.Sprintf("things like '%s %s' are impossible", invalidStepErrors.cmd, invalidStepErrors.obj)
		}

		var notEnoughObjectsErrors notEnoughObjectsError
		if errors.As(gameOverErrors.err, &notEnoughObjectsErrors) {
			return fmt.Sprintf("be careful with scarce %ss", notEnoughObjectsErrors.obj.name)
		}

		var commandLimitExceededErrors commandLimitExceededError
		if errors.As(gameOverErrors.err, &commandLimitExceededErrors) {
			return fmt.Sprintf("%s less", string(commandLimitExceededErrors.cmd)) //
		}

		var objectLimitExceededErrors objectLimitExceededError
		if errors.As(gameOverErrors.err, &objectLimitExceededErrors) {
			return fmt.Sprintf("don't be greedy, %d %s is enough", objectLimitExceededErrors.limit, objectLimitExceededErrors.obj) //
		}
	}

	// var invalidStepErrors invalidStepError
	// if errors.As(err, &invalidStepErrors) {
	// 	return fmt.Sprintf("things like %s %s are impossible", invalidStepErrors.cmd, invalidStepErrors.obj)
	// }

	// var notEnoughObjectsErrors notEnoughObjectsError
	// if errors.As(err, &notEnoughObjectsErrors) {
	// 	return fmt.Sprintf("be careful with scarce %ss", notEnoughObjectsErrors.obj.name)
	// }

	// var commandLimitExceededErrors commandLimitExceededError
	// if errors.As(err, &commandLimitExceededErrors) {
	// 	return fmt.Sprintf("%s less", string(commandLimitExceededErrors.cmd)) //
	// }

	// var objectLimitExceededErrors objectLimitExceededError
	// if errors.As(err, &objectLimitExceededErrors) {
	// 	return fmt.Sprintf("don't be greedy, %d %s is enough", objectLimitExceededErrors.limit, objectLimitExceededErrors.obj) //
	// }

	return "unknown error"

}

// конец решения

func main() {
	gm := newGame()
	steps := []step{
		{eat, apple},
		{talk, bob},
		{take, coin},
		{eat, mushroom},
	}
	// fmt.Println(string(eat))
	// fmt.Println(string(talk))
	for _, st := range steps {
		if err := tryStep(gm, st); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fmt.Println("You win!")
}

// tryStep() выполняет шаг игры и печатает результат
func tryStep(gm *game, st step) error {
	fmt.Printf("trying to %s %s... ", st.cmd, st.obj.name)
	if err := gm.execute(st); err != nil {
		fmt.Println("FAIL")
		return err
	}
	fmt.Println("OK")
	return nil
}
