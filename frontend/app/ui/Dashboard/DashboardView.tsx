'use client';

import { useState, useEffect, useCallback } from 'react';
import { Wallet, BarChart3, AlertTriangle, CheckCircle2, XCircle, AlertCircle } from 'lucide-react';
import { 
  DashboardLayout, 
  Dropzone, 
  DataTableCompound,
  TableColumn,
  CardSkeleton,
  Alert 
} from '@/app/components';
import { useNotification } from '@/app/contexts';
import { Transaction, mockApiCall, ApiResponse } from './mockApi';
import styles from './DashboardView.module.css';

export function DashboardView() {
  const { showSuccess, showError, showWarning } = useNotification();
  const [transactions, setTransactions] = useState<Transaction[]>([]);
  const [loading, setLoading] = useState(true);
  const [searchQuery, setSearchQuery] = useState('');
  const [currentPage, setCurrentPage] = useState(1);
  const [perPage] = useState(10);
  const [sortBy, setSortBy] = useState<string>('timestamp');
  const [sortOrder, setSortOrder] = useState<'asc' | 'desc'>('desc');
  const [apiResponse, setApiResponse] = useState<ApiResponse<Transaction> | null>(null);
  const [uploadError, setUploadError] = useState<string | null>(null);

  // Fetch data function (simulates server-side call)
  const fetchTransactions = useCallback(async () => {
    setLoading(true);
    try {
      const response = await mockApiCall<Transaction>({
        page: currentPage,
        per_page: perPage,
        search: searchQuery,
        sort_by: sortBy,
        order: sortOrder
      });
      
      if ('data' in response) {
        setApiResponse(response);
        setTransactions(response.data);
      }
    } catch (error) {
      console.error('Failed to fetch transactions:', error);
    } finally {
      setLoading(false);
    }
  }, [currentPage, perPage, searchQuery, sortBy, sortOrder]);

  useEffect(() => {
    fetchTransactions();
  }, [fetchTransactions]);

  const handleFileUpload = async (file: File) => {
    setUploadError(null);
    
    // Validate file type
    if (!file.name.endsWith('.csv')) {
      const errorMsg = 'Invalid file type. Please upload a CSV file.';
      setUploadError(errorMsg);
      showError(errorMsg, 'Upload Failed');
      return;
    }

    try {
      // Simulate upload process
      showWarning('Uploading file...', 'Processing');
      
      // TODO: Parse CSV and send to backend
      await new Promise(resolve => setTimeout(resolve, 1500));
      
      showSuccess(`${file.name} uploaded successfully!`, 'Upload Complete');
      setUploadError(null);
      
      // Refresh data after successful upload
      fetchTransactions();
    } catch (error) {
      const errorMsg = 'Failed to upload file. Please try again.';
      setUploadError(errorMsg);
      showError(errorMsg, 'Upload Failed');
    }
  };

  const handleSearch = (value: string) => {
    setSearchQuery(value);
    setCurrentPage(1); // Reset to first page on search
  };

  const handleSort = (key: string) => {
    if (sortBy === key) {
      setSortOrder(sortOrder === 'asc' ? 'desc' : 'asc');
    } else {
      setSortBy(key);
      setSortOrder('asc');
    }
    setCurrentPage(1);
  };

  const handlePageChange = (page: number) => {
    setCurrentPage(page);
  };

  // Calculate stats
  const calculateBalance = () => {
    return transactions
      .filter(t => t.status === 'SUCCESS')
      .reduce((acc, t) => {
        return t.type === 'CREDIT' ? acc + t.amount : acc - t.amount;
      }, 0);
  };

  const issues = transactions.filter(t => t.status === 'FAILED' || t.status === 'PENDING');

  const formatCurrency = (amount: number) => {
    return new Intl.NumberFormat('id-ID', {
      style: 'currency',
      currency: 'IDR',
      minimumFractionDigits: 0,
    }).format(amount);
  };

  const formatDate = (timestamp: number) => {
    return new Date(timestamp * 1000).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    });
  };

  const getStatusBadge = (status: string) => {
    const statusConfig = {
      SUCCESS: { icon: CheckCircle2, color: styles.statusSuccess },
      FAILED: { icon: XCircle, color: styles.statusFailed },
      PENDING: { icon: AlertCircle, color: styles.statusPending }
    };

    const config = statusConfig[status as keyof typeof statusConfig];
    if (!config) return null;

    const Icon = config.icon;
    return (
      <span className={`${styles.statusBadge} ${config.color}`}>
        <Icon size={14} />
        {status}
      </span>
    );
  };

  // Table columns definition
  const columns: TableColumn<Transaction>[] = [
    {
      key: 'timestamp',
      header: 'Date',
      sortable: true,
      render: (row: Transaction) => <span className={styles.dateCell}>{formatDate(row.timestamp)}</span>
    },
    {
      key: 'name',
      header: 'Name',
      sortable: true,
      render: (row: Transaction) => <span className={styles.nameCell}>{row.name}</span>
    },
    {
      key: 'type',
      header: 'Type',
      sortable: true,
      render: (row: Transaction) => (
        <span className={`${styles.typeBadge} ${row.type === 'CREDIT' ? styles.badgeCredit : styles.badgeDebit}`}>
          {row.type}
        </span>
      )
    },
    {
      key: 'amount',
      header: 'Amount',
      sortable: true,
      render: (row: Transaction) => (
        <span className={`${styles.amountCell} ${row.type === 'CREDIT' ? styles.amountCredit : styles.amountDebit}`}>
          {row.type === 'CREDIT' ? '+' : '-'}{formatCurrency(row.amount)}
        </span>
      )
    },
    {
      key: 'status',
      header: 'Status',
      sortable: true,
      render: (row: Transaction) => getStatusBadge(row.status)
    },
    {
      key: 'description',
      header: 'Description',
      render: (row: Transaction) => <span className={styles.descriptionCell}>{row.description}</span>
    }
  ];

  return (
    <DashboardLayout>
      <div className={styles.dashboard}>
        {/* Stats Cards */}
        <div className={styles.statsGrid}>
          {loading && !apiResponse ? (
            <>
              <CardSkeleton />
              <CardSkeleton />
              <CardSkeleton />
            </>
          ) : (
            <>
              <div className={styles.statCard}>
                <div className={styles.statIcon}>
                  <Wallet size={32} />
                </div>
                <div className={styles.statContent}>
                  <div className={styles.statLabel}>Current Balance</div>
                  <div className={styles.statValue}>{formatCurrency(calculateBalance())}</div>
                </div>
              </div>
              <div className={styles.statCard}>
                <div className={styles.statIcon}>
                  <BarChart3 size={32} />
                </div>
                <div className={styles.statContent}>
                  <div className={styles.statLabel}>Total Transactions</div>
                  <div className={styles.statValue}>{apiResponse?.meta.pagination.total || 0}</div>
                </div>
              </div>
              <div className={styles.statCard}>
                <div className={styles.statIcon}>
                  <AlertTriangle size={32} />
                </div>
                <div className={styles.statContent}>
                  <div className={styles.statLabel}>Issues</div>
                  <div className={styles.statValue}>{issues.length}</div>
                </div>
              </div>
            </>
          )}
        </div>

        {/* Upload Section */}
        <div className={styles.uploadContainer}>
          <Dropzone 
            onFileSelect={handleFileUpload}
            title="Upload Bank Statement"
            subtitle="Click or drag and drop your CSV file"
          />
          {uploadError && (
            <div style={{ marginTop: 'var(--spacing-md)' }}>
              <Alert 
                type="error" 
                message={uploadError}
                dismissible
                onClose={() => setUploadError(null)}
              />
            </div>
          )}
        </div>

        {/* All Transactions Table */}
        <DataTableCompound
          title="All Transactions"
          subtitle={apiResponse ? `Showing ${apiResponse.meta.pagination.count} of ${apiResponse.meta.pagination.total} transactions` : undefined}
          columns={columns}
          data={transactions}
          meta={apiResponse?.meta.pagination}
          loading={loading}
          searchValue={searchQuery}
          onSearchChange={handleSearch}
          onPageChange={handlePageChange}
          onSort={handleSort}
          sortBy={sortBy}
          sortOrder={sortOrder}
          showSearch={true}
          showPagination={true}
        />
      </div>
    </DashboardLayout>
  );
}

