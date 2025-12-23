import Axios, { URL } from "@/axios";
import { GetDataSuccess, type GetDataSuccessType } from "@/zod/get-data";
import { useQuery } from "@tanstack/react-query";

async function getData(pageIndex: number, pageSize: number, raw: boolean) {
  const response = await Axios.get(URL.GetData, {
    params: { offset: pageIndex * pageSize, limit: pageSize, raw },
  });

  return await GetDataSuccess.parseAsync(response.data);
}

export default function useGetData(
  pageIndex: number,
  pageSize: number,
  raw: boolean
) {
  const { data, isError, isLoading, refetch, isFetching } =
    useQuery<GetDataSuccessType>({
      queryFn: () => getData(pageIndex, pageSize, raw),
      queryKey: ["get-data", pageIndex, pageSize, raw],
      retry: 2,
    });

  return { data, isError, isLoading, refetch, isFetching };
}
