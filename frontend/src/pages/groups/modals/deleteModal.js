import Button from "../../../components/button";
import styles from "../../shared/style/modal.module.css";
import { useDeleteGroupMutation } from "../api/useDeleteGroupMutation";

export function DeleteModal({ deleteGroup, isOpen, original }) {
  const deleteGroupMutation = useDeleteGroupMutation();
  const id = original.id;

  if (!isOpen) {
    return null;
  }
  function deleteGroupHandle() {
    deleteGroupMutation.mutateAsync({ id });
    deleteGroup(true);
  }

  return (
    <div className={styles["backdrop"]}>
      <div className={styles["modal"]}>
        <p className={styles['delete-text']}>Вы уверены, что хотите удалить {original.name}?</p>
        <div className={styles["modal-delete"]}>
          <Button
            styleFeature='close'
            type="submit"
            onClick={() => deleteGroupHandle()}
          >
            Удалить
          </Button>
          <Button onClick={deleteGroup}>Отмена</Button>
        </div>
      </div>
    </div>
  );
}
