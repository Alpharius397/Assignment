export type HookParams = {
  onSuccess: () => void;
  onError: (message: string) => void;
  onFailure: () => void;
};

export type data = {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  [key: string]: any;
};
