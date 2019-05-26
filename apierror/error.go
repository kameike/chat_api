package apierror

import (
	"fmt"
	"github.com/pkg/errors"
	"runtime"
	"strings"
)

// === ERRORS
func Error(message string, err error) ChatAPIError {
	if err == nil {
		err := fmt.Errorf("%s :(%s)", message, generateSigniture())
		return &chatAPIError{
			err:          err,
			errorMessage: message,
		}
	}

	return &chatAPIError{
		err:          err,
		errorMessage: message,
	}
}

func NewError(message string) ChatAPIError {
	err := fmt.Errorf("%s :(%s)", message, generateSigniture())
	return &chatAPIError{
		err:          err,
		errorMessage: message,
	}
}

func WrapMessage(message string, err ChatAPIError) ChatAPIError {
	return &chatAPIError{
		err:          err.Err(),
		errorMessage: message,
	}
}

func generateSigniture() string {
	_, file, line, _ := runtime.Caller(2)
	file = strings.ReplaceAll(file, ".go", "")
	return fmt.Sprintf("%s:%d", file, line)
}

const SERVICE_DOWN = "現在メッセージングサービスをご利用できません。時間をあけて再度ご利用ください"
const FATAL_ERROR = "現在メッセージングサービスをご利用できません。"
const LOGIN_FAIL = "チャットサービスへのログインに失敗しました。"
const INVALID_USER = "有効なユーザーが存在しません。"
const FAILD_GET_MESSAGES = "メッセージの取得に失敗しました"

const POST_MESSAGE_FAILD = "メッセージの送信に失敗しました。"
const CHATROOM_NOT_FOUND = "チャットルームが見つかりません。"
const RELOAD_AFTER_POST_FALID = "メッセージは送信されましたが、更新に失敗しました。"
const FAILD_UPDATE_USER_INFO = "ユーザー情報の更新に失敗しました。"

const THINK_LATER = "現在メッセージングサービスをご利用できません。"
const INVALID_POST_MESSAGE = "不正なチャットルームにメッセージを送信しようとしています."
const GET_MESSAGE_FAIL = "メッセージの取得に失敗しました"

type ChatAPIError interface {
	Error() string
	Err() error
	ErrorMessage() string
	Wrap(error) ChatAPIError
	WrapWithSelf(wrapedTarget ChatAPIError) ChatAPIError
}

type chatAPIError struct {
	err          error
	errorMessage string
}

func (c *chatAPIError) Err() error {
	return c.err
}

func (c *chatAPIError) ErrorMessage() string {
	return c.errorMessage
}

func (c *chatAPIError) WrapWithSelf(wrapedTarget ChatAPIError) ChatAPIError {
	if wrapedTarget != nil {
		return &chatAPIError{
			err:          errors.Wrap(wrapedTarget.Err(), c.Error()),
			errorMessage: c.ErrorMessage(),
		}
	}
	return c
}
func (c *chatAPIError) Wrap(err error) ChatAPIError {
	if err != nil {
		c.err = errors.Wrap(err, err.Error())
	}
	return c
}

func (e *chatAPIError) Error() string {
	sysMsg := ""
	if e.err != nil {
		sysMsg = e.err.Error()
	}

	return fmt.Sprintf("ERROR %s: %s", e.errorMessage, sysMsg)
}
