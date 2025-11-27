import { testIndexApi } from './api';

// 页面加载完成后执行
document.addEventListener('DOMContentLoaded', () => {
  // 获取DOM元素
  const testApiBtn = document.getElementById('testApiBtn');
  const apiResponse = document.getElementById('apiResponse');

  if (testApiBtn && apiResponse) {
    // 添加按钮点击事件
    testApiBtn.addEventListener('click', async () => {
      apiResponse.textContent = '正在请求后端接口...';
      apiResponse.style.color = '#666';

      try {
        // 调用API测试函数
        const result = await testIndexApi();
        apiResponse.textContent = `后端响应: ${JSON.stringify(result, null, 2)}`;
        apiResponse.style.color = '#27ae60';
        console.log('API测试成功:', result);
      } catch (error) {
        apiResponse.textContent = `请求失败: ${error instanceof Error ? error.message : '未知错误'}`;
        apiResponse.style.color = '#e74c3c';
        console.error('API测试失败:', error);
      }
    });
  }
});

// 应用初始化函数
function initApp() {
  console.log('Go Blog前端应用初始化完成');
}

// 调用初始化函数
initApp();