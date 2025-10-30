import { DashboardLayout } from '@/app/components';
import { UnderMaintenance } from '@/app/ui/UnderMaintenance';

export default function SettingsPage() {
  return (
    <DashboardLayout>
      <UnderMaintenance 
        title="Settings Page"
        message="The settings page is currently under development. Please check back later."
      />
    </DashboardLayout>
  );
}

