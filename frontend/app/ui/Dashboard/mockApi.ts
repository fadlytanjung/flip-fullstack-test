// Mock Transaction Type
export interface Transaction {
  id: string;
  timestamp: number;
  name: string;
  type: 'CREDIT' | 'DEBIT';
  amount: number;
  status: 'SUCCESS' | 'FAILED' | 'PENDING';
  description: string;
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
  message: string;
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

// Mock data
const mockTransactions: Transaction[] = [
  { id: '1', timestamp: 1624507883, name: 'JOHN DOE', type: 'DEBIT', amount: 250000, status: 'SUCCESS', description: 'restaurant' },
  { id: '2', timestamp: 1624608050, name: 'E-COMMERCE A', type: 'DEBIT', amount: 150000, status: 'FAILED', description: 'clothes' },
  { id: '3', timestamp: 1624512883, name: 'COMPANY A', type: 'CREDIT', amount: 12000000, status: 'SUCCESS', description: 'salary' },
  { id: '4', timestamp: 1624615065, name: 'E-COMMERCE B', type: 'DEBIT', amount: 150000, status: 'PENDING', description: 'clothes' },
  { id: '5', timestamp: 1624520000, name: 'UTILITY COMPANY', type: 'DEBIT', amount: 500000, status: 'SUCCESS', description: 'electricity' },
  { id: '6', timestamp: 1624530000, name: 'GROCERY STORE', type: 'DEBIT', amount: 350000, status: 'SUCCESS', description: 'groceries' },
  { id: '7', timestamp: 1624540000, name: 'FREELANCE CLIENT', type: 'CREDIT', amount: 5000000, status: 'SUCCESS', description: 'project payment' },
  { id: '8', timestamp: 1624550000, name: 'ONLINE SHOP C', type: 'DEBIT', amount: 200000, status: 'PENDING', description: 'electronics' },
  { id: '9', timestamp: 1624560000, name: 'GAS STATION', type: 'DEBIT', amount: 300000, status: 'SUCCESS', description: 'fuel' },
  { id: '10', timestamp: 1624570000, name: 'RESTAURANT B', type: 'DEBIT', amount: 180000, status: 'SUCCESS', description: 'dining' },
  { id: '11', timestamp: 1624580000, name: 'PHARMACY', type: 'DEBIT', amount: 120000, status: 'FAILED', description: 'medicine' },
  { id: '12', timestamp: 1624590000, name: 'DIVIDEND PAYMENT', type: 'CREDIT', amount: 2000000, status: 'SUCCESS', description: 'investment' },
  { id: '13', timestamp: 1624600000, name: 'COFFEE SHOP', type: 'DEBIT', amount: 50000, status: 'SUCCESS', description: 'coffee' },
  { id: '14', timestamp: 1624610000, name: 'BOOKSTORE', type: 'DEBIT', amount: 250000, status: 'SUCCESS', description: 'books' },
  { id: '15', timestamp: 1624620000, name: 'INSURANCE', type: 'DEBIT', amount: 1500000, status: 'SUCCESS', description: 'monthly premium' },
  { id: '16', timestamp: 1624630000, name: 'RIDE SHARING', type: 'DEBIT', amount: 75000, status: 'SUCCESS', description: 'transportation' },
  { id: '17', timestamp: 1624640000, name: 'STREAMING SERVICE', type: 'DEBIT', amount: 100000, status: 'SUCCESS', description: 'subscription' },
  { id: '18', timestamp: 1624650000, name: 'GYM MEMBERSHIP', type: 'DEBIT', amount: 400000, status: 'PENDING', description: 'fitness' },
  { id: '19', timestamp: 1624660000, name: 'MOBILE RECHARGE', type: 'DEBIT', amount: 50000, status: 'SUCCESS', description: 'phone credit' },
  { id: '20', timestamp: 1624670000, name: 'CONSULTING FEE', type: 'CREDIT', amount: 3000000, status: 'SUCCESS', description: 'professional services' },
];

// Mock API call function
export async function mockApiCall<T>(params: {
  page?: number;
  per_page?: number;
  search?: string;
  sort_by?: string;
  order?: 'asc' | 'desc';
  filters?: Record<string, any>;
}): Promise<ApiResponse<T> | ApiError> {
  // Simulate network delay
  await new Promise(resolve => setTimeout(resolve, 500));

  const {
    page = 1,
    per_page = 10,
    search = '',
    sort_by = 'timestamp',
    order = 'desc',
    filters = {}
  } = params;

  // Validate parameters
  if (page < 1) {
    return {
      status: 400,
      error: 'InvalidQueryParameter',
      message: 'One or more query parameters are invalid.',
      details: {
        page: 'Page number must be a positive integer.'
      }
    };
  }

  if (per_page > 100) {
    return {
      status: 400,
      error: 'InvalidQueryParameter',
      message: 'One or more query parameters are invalid.',
      details: {
        per_page: 'Maximum per_page limit is 100.'
      }
    };
  }

  // Filter data
  let filteredData = [...mockTransactions];

  // Search filter
  if (search) {
    const searchLower = search.toLowerCase();
    filteredData = filteredData.filter(item =>
      Object.values(item).some(value =>
        String(value).toLowerCase().includes(searchLower)
      )
    );
  }

  // Apply additional filters
  Object.entries(filters).forEach(([key, value]) => {
    if (value) {
      filteredData = filteredData.filter(item => item[key as keyof typeof item] === value);
    }
  });

  // Sort data
  filteredData.sort((a, b) => {
    const aValue = a[sort_by as keyof typeof a];
    const bValue = b[sort_by as keyof typeof b];
    
    if (aValue < bValue) return order === 'asc' ? -1 : 1;
    if (aValue > bValue) return order === 'asc' ? 1 : -1;
    return 0;
  });

  // Pagination
  const total = filteredData.length;
  const total_pages = Math.ceil(total / per_page);
  const start = (page - 1) * per_page;
  const end = start + per_page;
  const paginatedData = filteredData.slice(start, end);

  // Build response
  const response: ApiResponse<T> = {
    message: 'Successfully retrieved paginated data list.',
    data: paginatedData as T[],
    meta: {
      pagination: {
        total,
        count: paginatedData.length,
        per_page,
        current_page: page,
        total_pages,
        links: {
          next: page < total_pages ? `/api/data?page=${page + 1}&per_page=${per_page}` : null,
          prev: page > 1 ? `/api/data?page=${page - 1}&per_page=${per_page}` : null
        }
      },
      filters,
      sort: {
        by: sort_by,
        order
      }
    }
  };

  return response;
}

