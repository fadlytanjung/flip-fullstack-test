'use client';

import { useState, useCallback } from 'react';

export interface UseLoadingReturn {
  // State
  tableLoading: boolean;
  statsLoading: boolean;
  uploading: boolean;
  
  // Actions
  setTableLoading: (loading: boolean) => void;
  setStatsLoading: (loading: boolean) => void;
  setUploading: (uploading: boolean) => void;
  
  // Helpers
  isAnyLoading: () => boolean;
}

export function useLoading(): UseLoadingReturn {
  const [tableLoading, setTableLoadingState] = useState(true);
  const [statsLoading, setStatsLoadingState] = useState(true);
  const [uploading, setUploadingState] = useState(false);

  // Memoize setters to prevent dependency array issues
  const setTableLoading = useCallback((loading: boolean) => {
    setTableLoadingState(loading);
  }, []);

  const setStatsLoading = useCallback((loading: boolean) => {
    setStatsLoadingState(loading);
  }, []);

  const setUploading = useCallback((uploading: boolean) => {
    setUploadingState(uploading);
  }, []);

  // Check if any loading is in progress
  const isAnyLoading = useCallback(() => {
    return tableLoading || statsLoading || uploading;
  }, [tableLoading, statsLoading, uploading]);

  return {
    tableLoading,
    statsLoading,
    uploading,
    setTableLoading,
    setStatsLoading,
    setUploading,
    isAnyLoading,
  };
}
