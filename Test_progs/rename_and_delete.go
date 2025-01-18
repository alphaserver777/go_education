package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Укажите путь к папке, в которой нужно переименовать файлы и папки
	dirPath := "D:/11. Программирование/"

	// Создаем слайс для хранения путей, которые нужно переименовать
	var pathsToRename []string

	// Сначала собираем все пути, которые нужно переименовать
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Проверяем, содержит ли имя "[SW.BAND]"
		if strings.Contains(info.Name(), "[SW.BAND]") {
			pathsToRename = append(pathsToRename, path)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Ошибка при обходе файлов и папок: %v\n", err)
		return
	}

	// Переименовываем файлы и папки в обратном порядке
	for i := len(pathsToRename) - 1; i >= 0; i-- {
		path := pathsToRename[i]
		info, err := os.Stat(path)
		if err != nil {
			fmt.Printf("Ошибка при получении информации о %s: %v\n", path, err)
			continue
		}

		// Получаем имя файла или папки
		name := info.Name()

		// Удаляем "[SW.BAND]" из имени
		newName := strings.Replace(name, "[SW.BAND] ", "", 1)

		// Получаем полный путь к новому имени
		newPath := filepath.Join(filepath.Dir(path), newName)

		// Переименовываем файл или папку
		err = os.Rename(path, newPath)
		if err != nil {
			fmt.Printf("Ошибка при переименовании %s: %v\n", path, err)
			continue
		}

		// Выводим сообщение о переименовании
		if info.IsDir() {
			fmt.Printf("Папка переименована: %s -> %s\n", name, newName)
		} else {
			fmt.Printf("Файл переименован: %s -> %s\n", name, newName)
		}
	}

	// Удаляем указанные файлы в каждой папке
	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Проверяем, что это файл
		if !info.IsDir() {
			// Список файлов для удаления
			filesToDelete := []string{
				"[DMC.RIP] Качай редкие курсы!.url",
				"[WWW.SW.BAND] 150000 курсов ждут тебя!.url",
				"[WWW.SW.BAND] Прочти перед изучением!.docx",
			}

			// Проверяем, является ли файл одним из тех, которые нужно удалить
			for _, fileToDelete := range filesToDelete {
				if info.Name() == fileToDelete {
					// Удаляем файл
					err := os.Remove(path)
					if err != nil {
						fmt.Printf("Ошибка при удалении файла %s: %v\n", path, err)
					} else {
						fmt.Printf("Файл удален: %s\n", path)
					}
					break
				}
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Ошибка при обходе файлов и папок для удаления: %v\n", err)
	}
}
