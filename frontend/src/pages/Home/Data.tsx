import { useMemo, useState } from "react";
import {
  flexRender,
  getCoreRowModel,
  getFacetedRowModel,
  getFacetedUniqueValues,
  getFilteredRowModel,
  getSortedRowModel,
  useReactTable,
  type ColumnDef,
} from "@tanstack/react-table";
import type { UserDataType } from "@/zod/get-data";
import useGetData from "@/hooks/useGetData";

import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { DataTablePagination } from "@/components/data-pagination";
import Sidebar from "@/pages/Home/SideBar";
import Loading from "@/components/Loading";
import Error from "@/components/Error";
import { Label } from "@/components/ui/label";
import { Switch } from "@/components/ui/switch";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import useProfile from "@/hooks/useProfile";
import { Button } from "@/components/ui/button";
import { Spinner } from "@/components/ui/spinner";
import { SidebarTrigger } from "@/components/ui/sidebar";

function SwitchComponent({
  onChange,
  value,
}: {
  onChange: (status: boolean) => void;
  value: boolean;
}) {
  return (
    <div className="flex items-center space-x-2">
      <Switch id="airplane-mode" onCheckedChange={onChange} checked={value} />
      <Label htmlFor="airplane-mode">Raw Data</Label>
    </div>
  );
}

type SquareCardProps = {
  title: string;
  description: string;
  content?: string;
  isLoading: boolean;
};

function SquareCard({
  title,
  description,
  content,
  isLoading,
}: SquareCardProps) {
  return (
    <Card className="w-full">
      <CardHeader>
        <CardTitle>{title}</CardTitle>
        <CardDescription>{description}</CardDescription>
      </CardHeader>
      <CardContent>
        {isLoading ? (
          <Button disabled size="sm">
            <Spinner />
            Loading...
          </Button>
        ) : (
          <p>{content}</p>
        )}
      </CardContent>
    </Card>
  );
}

function ProfileCard() {
  const { data, isLoading } = useProfile();

  return (
    <div className="flex flex-row gap-4 justify-between items-center">
      <SquareCard
        title="Username"
        description="Username of signed-in user"
        content={data?.user_name}
        isLoading={isLoading}
      />
      <SquareCard
        title="Email"
        description="Email of signed-in user"
        content={data?.email}
        isLoading={isLoading}
      />
      <SquareCard
        title="Aadhar ID"
        description="Aadhar ID of signed-in user"
        content={data?.aadhar}
        isLoading={isLoading}
      />
    </div>
  );
}

function DataTable() {
  const columns = useMemo<ColumnDef<UserDataType>[]>(
    () => [
      { accessorKey: "id", header: "ID" },
      { accessorKey: "user_name", header: "User name" },
      { accessorKey: "email", header: "Email" },
      { accessorKey: "aadhar", header: "Aadhar" },
    ],
    []
  );

  const [rawData, setRawData] = useState<boolean>(true);

  const [pagination, setPagination] = useState({
    pageIndex: 0,
    pageSize: 25,
  });

  const { data, isLoading, isError } = useGetData(
    pagination.pageIndex,
    pagination.pageSize,
    rawData
  );

  const flatData = useMemo(() => {
    return data?.data || [];
  }, [data]);

  const table = useReactTable({
    data: flatData,
    columns,
    getCoreRowModel: getCoreRowModel(),
    getSortedRowModel: getSortedRowModel(),
    state: {
      pagination,
    },
    pageCount: data
      ? Math.max(1, Math.ceil(data.total / pagination.pageSize))
      : 1,
    getFilteredRowModel: getFilteredRowModel(),
    getFacetedRowModel: getFacetedRowModel(),
    getFacetedUniqueValues: getFacetedUniqueValues(),
    onPaginationChange: (updater) => {
      setPagination((old) =>
        typeof updater === "function" ? updater(old) : updater
      );
    },
    manualPagination: true,
  });

  if (isLoading)
    return (
      <div className="hidden h-full flex-1 flex-col gap-8 p-8 pt-4 md:flex text-left">
        <Loading />
      </div>
    );

  if (isError)
    return (
      <div className="hidden h-full flex-1 flex-col gap-8 p-8 pt-4 md:flex text-left items-center justify-center">
        <Error />
      </div>
    );

  return (
    <div className="hidden h-full flex-1 flex-col gap-8 p-8 pt-4 md:flex text-left">
      <div className="w-full">
        <h2 className="text-2xl font-semibold tracking-tight mb-2">
          User Profile
        </h2>
        <ProfileCard />
      </div>
      <div className="flex items-center justify-between gap-2">
        <div className="flex flex-col gap-1">
          <h2 className="text-2xl font-semibold tracking-tight">
            List of Aadhar Data
          </h2>
          <p className="text-muted-foreground">
            Here&apos;s a list of all users present.
          </p>
        </div>
        <SwitchComponent onChange={setRawData} value={rawData} />
      </div>
      <div className="flex flex-col gap-4 w-full">
        <div className="overflow-hidden rounded-md border">
          <Table>
            <TableHeader>
              {table.getHeaderGroups().map((headerGroup) => (
                <TableRow key={headerGroup.id}>
                  {headerGroup.headers.map((header) => {
                    return (
                      <TableHead key={header.id} colSpan={header.colSpan}>
                        {header.isPlaceholder
                          ? null
                          : flexRender(
                              header.column.columnDef.header,
                              header.getContext()
                            )}
                      </TableHead>
                    );
                  })}
                </TableRow>
              ))}
            </TableHeader>
            <TableBody>
              {table.getRowModel().rows?.length ? (
                table.getRowModel().rows.map((row) => (
                  <TableRow key={row.id}>
                    {row.getVisibleCells().map((cell) => (
                      <TableCell key={cell.id}>
                        {flexRender(
                          cell.column.columnDef.cell,
                          cell.getContext()
                        )}
                      </TableCell>
                    ))}
                  </TableRow>
                ))
              ) : (
                <TableRow>
                  <TableCell
                    colSpan={columns.length}
                    className="h-24 text-center"
                  >
                    No results.
                  </TableCell>
                </TableRow>
              )}
            </TableBody>
          </Table>
        </div>
        <DataTablePagination table={table} />
      </div>
    </div>
  );
}

export default function DataPage() {
  return (
    <div className="flex w-full m-0 relative">
      <Sidebar />
      <div className="m-2">
        <SidebarTrigger />
      </div>
      <DataTable />
    </div>
  );
}
