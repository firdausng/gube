export type NavigationItem = {
	folder: string;
	name: string;
	path: string;
};

export type Pod = {
	name: string
	namespace: string
	phase: string
	containers: any
	owner: any
	menu: {
		podName:string
		namespace: string
	}
}