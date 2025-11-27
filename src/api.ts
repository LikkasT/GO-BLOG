import axios from 'axios';

// 创建axios实例，配置基本URL和超时时间
const apiClient = axios.create({
  baseURL: '/api', // 使用相对路径，利用Vite代理
  timeout: 10000, // 请求超时时间
  headers: {
    'Content-Type': 'application/json'
  }
});

// 定义API响应类型
export interface ApiResponse {
  message: string;
  [key: string]: any; // 允许额外的属性
}

/**
 * 测试后端/index接口
 * @returns Promise<ApiResponse> - 后端响应数据
 */
export async function testIndexApi(): Promise<ApiResponse> {
  try {
    const response = await apiClient.get<ApiResponse>('/index');
    return response.data;
  } catch (error) {
    if (typeof error === 'object' && error !== null && 'message' in error) {
      throw new Error(`API请求失败: ${(error as any).message}`);
    } else {
      throw new Error(`未知错误: ${error instanceof Error ? error.message : '无法连接到后端'}`);
    }
  }
}

/**
 * 这里可以添加更多API函数，例如：
 * - 获取博客列表
 * - 获取单篇博客详情
 * - 创建博客
 * - 更新博客
 * - 删除博客
 */
// export async function getBlogList(): Promise<Blog[]> {
//   const response = await apiClient.get<Blog[]>('/api/blogs');
//   return response.data;
// }