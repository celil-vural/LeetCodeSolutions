/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public int val;
 *     public ListNode next;
 *     public ListNode(int val=0, ListNode next=null) {
 *         this.val = val;
 *         this.next = next;
 *     }
 * }
 */
public class Solution {
    public ListNode RemoveNthFromEnd(ListNode head, int n) {
        if (head == null || n <= 0) {
            return head;
        }
        ListNode dummy = new ListNode(0);
        dummy.next = head;
        ListNode fast = dummy;
        ListNode slow = dummy;
        // Fast işaretçisini n adım ilerletme
        for (int i = 0; i <= n; i++) {
            if (fast == null) {
                return head; // Liste boyutu n'den küçükse veya boşsa, başlangıç durumunu döndür
            }
            fast = fast.next;
        }
        // Slow ve fast işaretçileri birlikte ilerletme
        while (fast != null) {
            slow = slow.next;
            fast = fast.next;
        }
        // Sondan n. elemanı çıkartma
        slow.next = slow.next.next;
        return dummy.next;
    }
}
