'use client';

import { useState, useCallback, useEffect } from 'react';
import { fetchTransactions, Transaction, ApiError } from '../Dashboard/api';

export interface UseTableProps {
  perPage?: number;
  sortBy?: string;
  sortOrder?: 'asc' | 'desc' | null;
}

export interface UseTableReturn {
  // State
  transactions: Transaction[];
  loading: boolean;
  searchQuery: string;
  currentPage: number;
  perPage: number;
  sortBy: string | null;
  sortOrder: 'asc' | 'desc' | null;
  apiResponse: any;
  filterType: string;
  filterStatus: string;
  
  // Actions
  handleSearch: (value: string) => void;
  handleSort: (key: string) => void;
  handlePageChange: (page: number) => void;
  handleFilterType: (type: string) => void;
  handleFilterStatus: (status: string) => void;
  refreshData: () => Promise<void>;
  
  // Formatters
  formatCurrency: (amount: number) => string;
  formatDate: (timestamp: number) => string;
}

export function useTable(props?: UseTableProps): UseTableReturn {
  const perPageDefault = props?.perPage || 10;
  const sortByDefault = props?.sortBy || null;
  const sortOrderDefault = props?.sortOrder || null;
  
  // State
  const [transactions, setTransactions] = useState<Transaction[]>([]);
  const [loading, setLoading] = useState(true);
  const [searchQuery, setSearchQuery] = useState('');
  const [currentPage, setCurrentPage] = useState(1);
  const [perPage] = useState(perPageDefault);
  const [sortBy, setSortBy] = useState<string | null>(sortByDefault);
  const [sortOrder, setSortOrder] = useState<'asc' | 'desc' | null>(sortOrderDefault);
  const [apiResponse, setApiResponse] = useState<any>(null);
  const [filterType, setFilterType] = useState<string>('');
  const [filterStatus, setFilterStatus] = useState<string>('');

  // Fetch transaction data
  const fetchTransactionData = useCallback(async (
    page: number,
    pageSize: number,
    search: string,
    sortField: string | null,
    sortDir: 'asc' | 'desc' | null,
    type?: string,
    status?: string
  ) => {
    setLoading(true);
    try {
      const response = await fetchTransactions({
        page,
        per_page: pageSize,
        search,
        sort_by: sortField || undefined,
        order: sortDir || undefined,
        type: type || undefined,
        status: status || undefined
      });
      
      if ('data' in response && !('error' in response)) {
        setApiResponse(response);
        setTransactions(response.data);
      } else if ('error' in response) {
        const error = response as ApiError;
        console.error('Failed to fetch transactions:', error.message);
      }
    } catch (error) {
      console.error('Failed to fetch transactions:', error);
    } finally {
      setLoading(false);
    }
  }, []);

  // Initial data load and refresh on filters change
  useEffect(() => {
    fetchTransactionData(currentPage, perPage, searchQuery, sortBy, sortOrder, filterType, filterStatus);
  }, [currentPage, perPage, searchQuery, sortBy, sortOrder, filterType, filterStatus, fetchTransactionData]);

  // Handle search
  const handleSearch = useCallback((value: string) => {
    setSearchQuery(value);
    setCurrentPage(1);
  }, []);

  // Handle sort (three-state: NULL -> ASC -> DESC -> NULL)
  const handleSort = useCallback((key: string) => {
    setSortBy((prevSortBy) => {
      setSortOrder((prevSortOrder) => {
        // If clicking a different column, start with ASC
        if (prevSortBy !== key) {
          return 'asc';
        }
        
        // Same column - cycle through states
        if (prevSortOrder === null) {
          // NULL -> ASC
          return 'asc';
        } else if (prevSortOrder === 'asc') {
          // ASC -> DESC
          return 'desc';
        } else {
          // DESC -> NULL (reset)
          return null;
        }
      });
      
      // Update sortBy only if it's a different column
      return prevSortBy === key ? prevSortBy : key;
    });
    
    setCurrentPage(1);
  }, []);

  // Handle page change
  const handlePageChange = useCallback((page: number) => {
    setCurrentPage(page);
  }, []);

  // Handle filter type
  const handleFilterType = useCallback((type: string) => {
    setFilterType(type);
    setCurrentPage(1);
  }, []);

  // Handle filter status
  const handleFilterStatus = useCallback((status: string) => {
    setFilterStatus(status);
    setCurrentPage(1);
  }, []);

  // Refresh data manually
  const refreshData = useCallback(async () => {
    await fetchTransactionData(currentPage, perPage, searchQuery, sortBy, sortOrder, filterType, filterStatus);
  }, [currentPage, perPage, searchQuery, sortBy, sortOrder, filterType, filterStatus, fetchTransactionData]);

  // Format currency
  const formatCurrency = useCallback((amount: number) => {
    return new Intl.NumberFormat('id-ID', {
      style: 'currency',
      currency: 'IDR',
      minimumFractionDigits: 0,
    }).format(amount);
  }, []);

  // Format date
  const formatDate = useCallback((timestamp: number) => {
    return new Date(timestamp * 1000).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    });
  }, []);

  return {
    transactions,
    loading,
    searchQuery,
    currentPage,
    perPage,
    sortBy,
    sortOrder,
    apiResponse,
    filterType,
    filterStatus,
    handleSearch,
    handleSort,
    handlePageChange,
    handleFilterType,
    handleFilterStatus,
    refreshData,
    formatCurrency,
    formatDate,
  };
}
