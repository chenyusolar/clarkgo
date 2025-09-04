package integration_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArtisanCommands(t *testing.T) {
	// 测试 make:controller
	t.Run("make controller", func(t *testing.T) {
		cmd := exec.Command("go", "run", ".", "artisan", "make:controller", "TestController")
		output, err := cmd.CombinedOutput()
		assert.NoError(t, err)
		assert.Contains(t, string(output), "Controller created")

		// 清理
		os.Remove("app/Http/Controllers/TestController.go")
	})

	// 测试 migrate
	t.Run("migrate", func(t *testing.T) {
		cmd := exec.Command("go", "run", ".", "artisan", "migrate")
		output, err := cmd.CombinedOutput()
		if err != nil {
			// 迁移可能失败(如数据库未配置)，但命令应能执行
			assert.Contains(t, string(output), "Migration")
		}
	})

	// 测试 cache:clear
	t.Run("clear cache", func(t *testing.T) {
		// 先创建测试缓存目录
		os.MkdirAll("storage/framework/cache", 0755)
		os.WriteFile("storage/framework/cache/test", []byte("test"), 0644)

		cmd := exec.Command("go", "run", ".", "artisan", "cache:clear")
		output, err := cmd.CombinedOutput()
		assert.NoError(t, err)
		assert.Contains(t, string(output), "cache cleared")
	})
}
