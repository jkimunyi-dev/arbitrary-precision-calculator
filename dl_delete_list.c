
#include "common.h"

int dl_delete_list(node **head, node **tail)
{
        /* If list is empty */
        if (*head == NULL || *tail == NULL)
                return FAILURE;

        /* If not */
        while ((*head) != NULL)
        {
                *head = (*head)->next;
                free(*head);
        }
        *head = NULL;
        *tail = NULL;
        return SUCCESS;
}
