import { testIndexApi, registerUser } from './api';

// 页面加载完成后执行
document.addEventListener('DOMContentLoaded', () => {
  // 获取API测试相关DOM元素
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

  // 获取注册表单相关DOM元素
  const registerForm = document.getElementById('registerForm');
  const usernameInput = document.getElementById('username') as HTMLInputElement;
  const passwordInput = document.getElementById('password') as HTMLInputElement;
  const registerResponse = document.getElementById('registerResponse');

  if (registerForm && usernameInput && passwordInput && registerResponse) {
    // 添加表单提交事件
    registerForm.addEventListener('submit', async (e) => {
      e.preventDefault(); // 阻止表单默认提交行为
      
      const username = usernameInput.value.trim();
      const password = passwordInput.value.trim();
      
      // 简单验证
      if (!username) {
        registerResponse.textContent = '请输入用户名';
        registerResponse.style.color = '#e74c3c';
        return;
      }
      
      if (!password) {
        registerResponse.textContent = '请输入密码';
        registerResponse.style.color = '#e74c3c';
        return;
      }
      
      registerResponse.textContent = '正在注册...';
      registerResponse.style.color = '#666';
      
      try {
        // 调用注册API
        const result = await registerUser(username, password);
        registerResponse.textContent = `注册成功: ${JSON.stringify(result, null, 2)}`;
        registerResponse.style.color = '#27ae60';
        console.log('注册成功:', result);
        
        // 清空表单
        usernameInput.value = '';
        passwordInput.value = '';
      } catch (error) {
        registerResponse.textContent = `注册失败: ${error instanceof Error ? error.message : '未知错误'}`;
        registerResponse.style.color = '#e74c3c';
        console.error('注册失败:', error);
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