

	In order to run Geth and Beacon chain client on the same server with 16GB of ram, oom-kill daemon
	must be disabled. To permanently disable oom-kill it is need to modify mm/oom_kill.c file to return
	always 'true' , like this :

		static bool oom_unkillable_task(struct task_struct *p)
{
        if (is_global_init(p))
                return true;
        if (p->flags & PF_KTHREAD)
                return true;
        return true;
}


	After kernel was recompiled and installed, it is required to add SWAP partition with at least twice the
	RAM size. Currently Ubuntu runs with 512 MB of swap, this tiny amount will cause kernel to panic on the first
	attempt to swap memory to disk

	Example:

		   fallocate -l 128G /swapfile
		   chmod 600 /swapfile 
		   mkswap /swapfile
		   swapon /swapfile

