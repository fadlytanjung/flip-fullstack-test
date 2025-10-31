// API Client Configuration
const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:9000/api';

// Transaction Type
export interface Transaction {
  id: string;
  timestamp: number;
  name: string;
  type: 'CREDIT' | 'DEBIT';
  amount: number;
  status: 'SUCCESS' | 'FAILED' | 'PENDING';
  description: string;
  created_at?: string;
}

// API Response Types
export interface PaginationLinks {
  next: string | null;
  prev: string | null;
}

export interface PaginationMeta {
  total: number;
  count: number;
  per_page: number;
  current_page: number;
  total_pages: number;
  links: PaginationLinks;
}

export interface ApiResponse<T> {
  message?: string;
  status?: number;
  data: T[];
  meta: {
    pagination: PaginationMeta;
    filters?: Record<string, any>;
    sort?: {
      by: string;
      order: string;
    };
  };
}

export interface ApiError {
  status: number;
  error: string;
  message: string;
  details?: Record<string, string>;
}

export interface UploadResponse {
  message: string;
  total_records: number;
  success_records: number;
  failed_records: number;
  pending_records: number;
}

// Fetch transactions with filtering, sorting, and pagination
export async function fetchTransactions(params: {
  page?: number;
  per_page?: number;
  search?: string;
  sort_by?: string;
  order?: 'asc' | 'desc';
  type?: string;
  status?: string;
}): Promise<ApiResponse<Transaction> | ApiError> {
  try {
    const queryParams = new URLSearchParams();
    
    if (params.page) queryParams.append('page', params.page.toString());
    if (params.per_page) queryParams.append('page_size', params.per_page.toString());
    if (params.search) queryParams.append('search', params.search);
    if (params.sort_by) queryParams.append('sort_by', params.sort_by);
    if (params.order) queryParams.append('sort_order', params.order.toUpperCase());
    if (params.type) queryParams.append('type', params.type);
    if (params.status) queryParams.append('status', params.status);

    const response = await fetch(`${API_BASE_URL}/transactions?${queryParams.toString()}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    });

    if (!response.ok) {
      const error = await response.json();
      return {
        status: response.status,
        error: error.error || 'Unknown error',
        message: error.message || 'Failed to fetch transactions',
        details: error.details,
      };
    }

    const data = await response.json();
    // Map backend response to frontend format
    return {
      message: data.data?.message || 'Successfully retrieved transactions',
      data: data.data?.data || [],
      meta: data.data?.meta || {
        pagination: {
          total: 0,
          count: 0,
          per_page: 10,
          current_page: 1,
          total_pages: 0,
          links: { next: null, prev: null },
        },
      },
    };
  } catch (error) {
    console.error('Failed to fetch transactions:', error);
    return {
      status: 500,
      error: 'NetworkError',
      message: error instanceof Error ? error.message : 'Failed to fetch transactions',
    };
  }
}

// Upload CSV file
export async function uploadFile(file: File): Promise<UploadResponse | ApiError> {
  try {
    const formData = new FormData();
    formData.append('file', file);

    const response = await fetch(`${API_BASE_URL}/upload`, {
      method: 'POST',
      body: formData,
    });

    if (!response.ok) {
      const error = await response.json();
      return {
        status: response.status,
        error: error.error || 'Unknown error',
        message: error.message || 'Failed to upload file',
        details: error.details,
      };
    }

    const data = await response.json();
    return data.data as UploadResponse;
  } catch (error) {
    console.error('Failed to upload file:', error);
    return {
      status: 500,
      error: 'NetworkError',
      message: error instanceof Error ? error.message : 'Failed to upload file',
    };
  }
}

// Get balance
export async function fetchBalance() {
  try {
    const response = await fetch(`${API_BASE_URL}/balance`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    });

    if (!response.ok) {
      const error = await response.json();
      return null;
    }

    const data = await response.json();
    return data.data;
  } catch (error) {
    console.error('Failed to fetch balance:', error);
    return null;
  }
}

// Fetch issues count
export async function fetchIssuesCount(): Promise<number> {
  try {
    const response = await fetch(`${API_BASE_URL}/issues?page=1&page_size=1`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
    });

    if (!response.ok) {
      return 0;
    }

    const data = await response.json();
    return data.data?.meta?.pagination?.total || 0;
  } catch (error) {
    console.error('Failed to fetch issues count:', error);
    return 0;
  }
}

// Health check
export async function healthCheck(): Promise<boolean> {
  try {
    const response = await fetch(`${API_BASE_URL}/health`, {
      method: 'GET',
    });
    return response.ok;
  } catch (error) {
    console.error('Health check failed:', error);
    return false;
  }
}

